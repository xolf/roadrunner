// Copyright (c) 2018 SpiralScout
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package http

import (
	"errors"
	"github.com/spf13/cobra"
	rr "github.com/spiral/roadrunner/cmd/rr/cmd"
	"github.com/spiral/roadrunner/service/rpc"
	"github.com/spiral/roadrunner/service"
	"github.com/spiral/roadrunner/cmd/rr/utils"
)

func init() {
	rr.CLI.AddCommand(&cobra.Command{
		Use:   "http:reset",
		Short: "Reload RoadRunner worker pools for the HTTP service",
		RunE:  reloadHandler,
	})
}

func reloadHandler(cmd *cobra.Command, args []string) error {
	svc, st := rr.Container.Get(rpc.Name)
	if st < service.StatusConfigured {
		return errors.New("RPC service is not configured")
	}

	client, err := svc.(*rpc.Service).Client()
	if err != nil {
		return err
	}
	defer client.Close()

	utils.Printf("<green>restarting http worker pool</reset>: ")

	var r string
	if err := client.Call("http.Reset", true, &r); err != nil {
		return err
	}

	utils.Printf("<green+hb>done</reset>")
	return nil
}
