//     Copyright (C) 2020, IrineSistiana
//
//     This file is part of mosdns.
//
//     mosdns is free software: you can redistribute it and/or modify
//     it under the terms of the GNU General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     mosdns is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

package pipeline

import (
	"context"
	"errors"
	"github.com/IrineSistiana/mosdns/dispatcher/handler"
	"github.com/IrineSistiana/mosdns/dispatcher/mlog"
	"github.com/sirupsen/logrus"
)

const PluginType = "pipeline"

func init() {
	handler.RegInitFunc(PluginType, Init)
}

var _ handler.ExecutablePlugin = (*pipelineRouter)(nil)

type pipelineRouter struct {
	tag    string
	logger *logrus.Entry

	args *Args
}

type Args struct {
	Pipe []string `yaml:"pipe"`
}

func Init(tag string, argsMap map[string]interface{}) (p handler.Plugin, err error) {
	args := new(Args)
	err = handler.WeakDecode(argsMap, args)
	if err != nil {
		return nil, handler.NewErrFromTemplate(handler.ETInvalidArgs, err)
	}

	if len(args.Pipe) == 0 {
		return nil, errors.New("empty pipeline")
	}

	return &pipelineRouter{
		tag:    tag,
		logger: mlog.NewPluginLogger(tag),
		args:   args,
	}, nil
}

func (pr *pipelineRouter) Tag() string {
	return pr.tag
}

func (pr *pipelineRouter) Type() string {
	return PluginType
}

func (pr *pipelineRouter) Exec(ctx context.Context, qCtx *handler.Context) (err error) {
	err = handler.NewPipeContext(pr.args.Pipe, pr.logger).ExecNextPlugin(ctx, qCtx)
	if err != nil {
		return handler.NewPluginError(pr.tag, err)
	}
	return nil
}
