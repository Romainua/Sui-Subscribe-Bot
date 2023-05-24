package text

var Example string
var Start string

func Init_text() {
	Example = `
List of filter subscribe example (currently on mainnet network):

Subscribe to all activity in network (too many tx)
{"jsonrpc":"2.0", "id": 1, "method": "suix_subscribeEvent", "params": [{"All":[]}]}

Subscribe to all activity from address
{"jsonrpc":"2.0", "id": 1, "method": "suix_subscribeEvent", "params": [{"All":[{"Sender": "0xb7847468db546ba85acb9dcdc0c5190b3ca6427d713ff52a4f8183c81f8a39e1"}]}]}

Suscribe to Staking Events
{"jsonrpc":"2.0", "id": 1, "method": "suix_subscribeEvent", "params": [{"MoveEventType": "0x3::validator::StakingRequestEvent"}]}

Suscribe to Staking Events by custom validator name
{"jsonrpc":"2.0", "id": 1, "method": "suix_subscribeEvent", "params": [{"And":[{"MoveEventField": {"path":"/validator_address","value":"0x1290ab8bca1c136d2b86967674a43a0f5cb30b352f90233cc37ebf7452b96787"}},{"MoveEventType": "0x3::validator::StakingRequestEvent"}]}]}

if you need any help to create your own filter feel free dming to @romanosadchyi Thanks!`

	Start = `
Hi, User! Just send json requsts via your filters like (currently on mainnet network):

     {"jsonrpc":"2.0", "id": 1, "method": "suix_subscribeEvent", "params": [{"All":[]}]}
	
Input /example to check more example.
More info are there https://docs.sui.io/build/event_api#event-filters`
}
