/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	pb "klook.com/protos/alarmemailpb"
	"klook.libs/krpc"
	"log"
	"service/erpdatacenterserv/conf"
	"time"
)

func main() {

	krpc.InitRPCClient(&conf.Conf.RPCConfig) // 每个服务只需要在main.go中 InitRPCClient 一次, 放在初始化配置InitConfig和InitCommParams之后

	c := pb.NewAlarmEmailHandleClient(krpc.TargetKlookErpDataCentersrv)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.AlarmEmailHandleFunc(ctx, &pb.AlarmEmailRequest{Subject: "title", Content: "content"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r)
}
