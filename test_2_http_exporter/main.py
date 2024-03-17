from flask import Flask, Response
import requests, os
from prometheus_client import Gauge, generate_latest, CollectorRegistry, CONTENT_TYPE_LATEST

app = Flask(__name__)
INFURA_ID = os.environ.get("INFURA_ID", "6d6abba470cc4091ba922c3d09459f51")

registry = CollectorRegistry()
METRICS = {
    'skymavis_ankr_infura_sync_lag': Gauge('skymavis_ankr_infura_sync_lag', 'Block number diff between infura and ankr', 
                                        registry=registry),
    'skymavis_ankr_infura_success': Gauge('skymavis_ankr_infura_success', 'Success if lag between infura and ankr is less than 5', 
                                        registry=registry),
}

def get_ankr_block_number():
    r = requests.post(
        "https://rpc.ankr.com/eth", 
        json = {"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}
    )
    return (int(r.json()["result"], 16))
    

def get_infura_block_number():
    r = requests.post(
        "https://mainnet.infura.io/v3/{}".format(INFURA_ID),
        json = {
            "jsonrpc": "2.0", 
            "method": "eth_blockNumber", 
            "params": [], 
            "id": 1
        }
    )
    return (int(r.json()["result"], 16))

@app.route('/metrics')
def metrics():
    ankr_infura_diff = get_ankr_block_number() - get_infura_block_number()
    METRICS["skymavis_ankr_infura_sync_lag"].set(ankr_infura_diff)
    METRICS["skymavis_ankr_infura_success"].set(
        1 if abs(ankr_infura_diff) < 5 else 0
    )
    
    return Response(generate_latest(registry), mimetype=CONTENT_TYPE_LATEST)

    
if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3000)