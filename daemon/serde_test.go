package daemon

import (
	"encoding/json"
	"testing"
)

func TestSerdeGetPeers(t *testing.T) {
	serialized := `{
        "peers": [
			{
                "addr": "162.19.249.100:2125",
                "connected_on": 1711663198,
                "cumulative_difficulty": "874788276435001",
                "height": 21939,
                "id": 7089875151156203202,
                "last_ping": 1711664680,
                "local_port": 2125,
                "peers": {
                    "255.255.255.255:2125": {
                        "out": {
                            "sent_at": 1746147705443
                        }
                    },
                    "74.208.251.149:2125": {
                        "both": {
                            "received_at": 1745996467029,
                            "sent_at": 1745953407894
                        }
                    },
                    "74.208.251.149:2126": {
                        "in": {
                            "received_at": 1745996467029
                        }
                    }
                },
                "pruned_topoheight": null,
                "tag": null,
                "top_block_hash": "0000000007eeed3fecdaedff82ad867a224826230c12465cf39186471e2e360e",
                "topoheight": 22241,
                "version": "1.8.0-58bb439"
            },
            {
                "addr": "74.208.251.149:2125",
                "connected_on": 1711663199,
                "cumulative_difficulty": "874788276435001",
                "height": 21939,
                "id": 2448648666414530279,
                "last_ping": 1711664682,
                "local_port": 2125,
                "peers": {},
                "pruned_topoheight": null,
                "tag": null,
                "top_block_hash": "0000000007eeed3fecdaedff82ad867a224826230c12465cf39186471e2e360e",
                "topoheight": 22241,
                "version": "1.8.0-58bb439"
            }
        ],
        "total_peers": 4,
        "hidden_peers": 0
    }`

	var peers GetPeersResult
	if err := json.Unmarshal([]byte(serialized), &peers); err != nil {
		t.Fatalf("failed to deserialize: %v", err)
	}

	_, err := json.Marshal(peers)
	if err != nil {
		t.Fatalf("failed to serialize: %v", err)
	}
}
