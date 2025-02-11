// Copyright 2022 Chainguard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package exec

import (
	"fmt"
	"log"
	"os/exec"
)

type Executor struct {
	WorkDir  string
	UseProot bool
	UseQemu  string
	Log      *log.Logger
}

type Option func(*Executor) error

func New(workDir string, logger *log.Logger, opts ...Option) (*Executor, error) {
	e := &Executor{
		WorkDir: workDir,
		Log:     logger,
	}

	for _, opt := range opts {
		if err := opt(e); err != nil {
			return nil, err
		}
	}

	return e, nil
}

func WithProot(proot bool) Option {
	return func(e *Executor) error {
		e.UseProot = proot
		return nil
	}
}

func WithQemu(qemuArch string) Option {
	return func(e *Executor) error {
		emu, err := exec.LookPath(fmt.Sprintf("qemu-%s", qemuArch))
		if err != nil {
			return fmt.Errorf("unable to find qemu emulator for %s: %w", qemuArch, err)
		}

		e.UseQemu = emu
		return nil
	}
}
