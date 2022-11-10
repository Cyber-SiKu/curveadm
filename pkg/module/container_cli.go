/*
 *  Copyright (c) 2021 NetEase Inc.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

/*
 * Project: CurveAdm
 * Created Date: 2021-12-09
 * Author: Jingli Chen (Wine93)
 */

// __SIGN_BY_WINE93__

package module

import (
	"fmt"
	"strings"
	"text/template"
)

const (
	TEMPLATE_INFO                = "{{.controller}} info"
	TEMPLATE_PULL_IMAGE          = "{{.controller}} pull {{.options}} {{.name}}"
	TEMPLATE_CREATE_CONTAINER    = "{{.controller}} create {{.options}} {{.image}} {{.command}}"
	TEMPLATE_START_CONTAINER     = "{{.controller}} start {{.options}} {{.containers}}"
	TEMPLATE_STOP_CONTAINER      = "{{.controller}} stop {{.options}} {{.containers}}"
	TEMPLATE_RESTART_CONTAINER   = "{{.controller}} restart {{.options}} {{.containers}}"
	TEMPLATE_WAIT_CONTAINER      = "{{.controller}} wait {{.options}} {{.containers}}"
	TEMPLATE_REMOVE_CONTAINER    = "{{.controller}} rm {{.options}} {{.containers}}"
	TEMPLATE_LIST_CONTAINERS     = "{{.controller}} ps {{.options}}"
	TEMPLATE_CONTAINER_EXEC      = "{{.controller}} exec {{.options}} {{.container}} {{.command}}"
	TEMPLATE_COPY_FROM_CONTAINER = "{{.controller}} cp {{.options}} {{.container}}:{{.srcPath}} {{.destPath}}"
	TEMPLATE_COPY_INTO_CONTAINER = "{{.controller}} cp {{.options}}  {{.srcPath}} {{.container}}:{{.destPath}}"
	TEMPLATE_INSPECT_CONTAINER   = "{{.controller}} inspect {{.options}} {{.container}}"
	TEMPLATE_CONTAINER_LOGS      = "{{.controller}} logs {{.options}} {{.container}}"
)

type ContainerCli struct {
	sshClient *SSHClient
	options   []string
	tmpl      *template.Template
	data      map[string]interface{}
}

func NewContainerCli(sshClient *SSHClient) *ContainerCli {
	cli := &ContainerCli{
		sshClient: sshClient,
		options:   []string{},
		tmpl:      nil,
		data:      map[string]interface{}{},
	}
	return cli
}

func (s *ContainerCli) AddOption(format string, args ...interface{}) *ContainerCli {
	s.options = append(s.options, fmt.Sprintf(format, args...))
	return s
}

func (cli *ContainerCli) Execute(options ExecOptions) (string, error) {
	cli.data["controller"] = options.ExecController
	cli.data["options"] = strings.Join(cli.options, " ")
	return execCommand(cli.sshClient, cli.tmpl, cli.data, options)
}

func (cli *ContainerCli) ContainerInfo() *ContainerCli {
	cli.tmpl = template.Must(template.New("ContainerInfo").Parse(TEMPLATE_INFO))
	return cli
}

func (cli *ContainerCli) PullImage(image string) *ContainerCli {
	cli.tmpl = template.Must(template.New("PullImage").Parse(TEMPLATE_PULL_IMAGE))
	cli.data["name"] = image
	return cli
}

func (cli *ContainerCli) CreateContainer(image, command string) *ContainerCli {
	cli.tmpl = template.Must(template.New("CreateContainer").Parse(TEMPLATE_CREATE_CONTAINER))
	cli.data["image"] = image
	cli.data["command"] = command
	return cli
}

func (cli *ContainerCli) StartContainer(containerId ...string) *ContainerCli {
	cli.tmpl = template.Must(template.New("StartContainer").Parse(TEMPLATE_START_CONTAINER))
	cli.data["containers"] = strings.Join(containerId, " ")
	return cli
}

func (cli *ContainerCli) StopContainer(containerId ...string) *ContainerCli {
	cli.tmpl = template.Must(template.New("StopContainer").Parse(TEMPLATE_STOP_CONTAINER))
	cli.data["containers"] = strings.Join(containerId, " ")
	return cli
}

func (cli *ContainerCli) RestartContainer(containerId ...string) *ContainerCli {
	cli.tmpl = template.Must(template.New("RestartContainer").Parse(TEMPLATE_RESTART_CONTAINER))
	cli.data["containers"] = strings.Join(containerId, " ")
	return cli
}

func (cli *ContainerCli) WaitContainer(containerId ...string) *ContainerCli {
	cli.tmpl = template.Must(template.New("WaitContainer").Parse(TEMPLATE_WAIT_CONTAINER))
	cli.data["containers"] = strings.Join(containerId, " ")
	return cli
}

func (cli *ContainerCli) RemoveContainer(containerId ...string) *ContainerCli {
	cli.tmpl = template.Must(template.New("RemoveContainer").Parse(TEMPLATE_REMOVE_CONTAINER))
	cli.data["containers"] = strings.Join(containerId, " ")
	return cli
}

func (cli *ContainerCli) ListContainers() *ContainerCli {
	cli.tmpl = template.Must(template.New("ListContainers").Parse(TEMPLATE_LIST_CONTAINERS))
	return cli
}

func (cli *ContainerCli) ContainerExec(containerId, command string) *ContainerCli {
	cli.tmpl = template.Must(template.New("ContainerExec").Parse(TEMPLATE_CONTAINER_EXEC))
	cli.data["container"] = containerId
	cli.data["command"] = command
	return cli
}

func (cli *ContainerCli) CopyFromContainer(containerId, srcPath, destPath string) *ContainerCli {
	cli.tmpl = template.Must(template.New("CopyFromContainer").Parse(TEMPLATE_COPY_FROM_CONTAINER))
	cli.data["container"] = containerId
	cli.data["srcPath"] = srcPath
	cli.data["destPath"] = destPath
	return cli
}

func (cli *ContainerCli) CopyIntoContainer(srcPath, containerId, destPath string) *ContainerCli {
	cli.tmpl = template.Must(template.New("CopyIntoContainer").Parse(TEMPLATE_COPY_INTO_CONTAINER))
	cli.data["srcPath"] = srcPath
	cli.data["container"] = containerId
	cli.data["destPath"] = destPath
	return cli
}

func (cli *ContainerCli) InspectContainer(containerId string) *ContainerCli {
	cli.tmpl = template.Must(template.New("InspectContainer").Parse(TEMPLATE_INSPECT_CONTAINER))
	cli.data["container"] = containerId
	return cli
}

func (cli *ContainerCli) ContainerLogs(containerId string) *ContainerCli {
	cli.tmpl = template.Must(template.New("ContainerLogs").Parse(TEMPLATE_CONTAINER_LOGS))
	cli.data["container"] = containerId
	return cli
}
