// Copyright (c) 2019 NEC Laboratories Europe GmbH.
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

package pb

import (
	"google.golang.org/protobuf/proto"

	"github.com/hyperledger-labs/minbft/messages"
	"github.com/hyperledger-labs/minbft/usig"
)

func MarshalOrPanic(m proto.Message) []byte {
	bytes, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	return bytes
}

func HelloFromAPI(h messages.Hello) *Hello {
	return &Hello{
		ReplicaId: h.ReplicaID(),
	}
}

func RequestFromAPI(req messages.Request) *Request {
	return &Request{
		ClientId:  req.ClientID(),
		Seq:       req.Sequence(),
		Operation: req.Operation(),
		Signature: req.Signature(),
	}
}

func PrepareFromAPI(prep messages.Prepare) *Prepare {
	return &Prepare{
		ReplicaId: prep.ReplicaID(),
		View:      prep.View(),
		Request:   RequestFromAPI(prep.Request()),
		Ui:        usig.MustMarshalUI(prep.UI()),
	}
}
