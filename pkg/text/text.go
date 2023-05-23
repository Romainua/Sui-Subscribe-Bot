package text

var Example string
var Start string

func Init_text() {
	Example = `
List of filter subscribe example (currently on devnet network):

Subscribe to Move Event via Package 0x2 and Module devnet_nft
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"MoveEvent"}, {"Package":"0x2"}, {"Module":"devnet_nft"}]}]}
  
Subscribe to only Publish activity
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"Publish"}]}]}
   
Subscribe to all activity in network (too many tx)
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[]}]}

Subscribe to all activity from address
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"SenderAddress": "0xc4173a804406a365e69dfb297d4eaaf002546ebd"}]}]}

Subscribe to activity from address which include only coin balance changing event
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"SenderAddress":"0xc4173a804406a365e69dfb297d4eaaf002546ebd"},{"EventType":"CoinBalanceChange"}]}]}

Suscribe to all activity which include Transfer Object
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"TransferObject"}]}]}

Suscribe to only Transfer Object activity from address
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"TransferObject"}, {"SenderAddress": "0xc4173a804406a365e69dfb297d4eaaf002546ebd"}]}]}

check activity of your package, just replace Package to your
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"MoveEvent"},{"Package":"0xc8832b8cddf246f3fc010c10f9d62d89c163d8cc"}]}]}

check activity of your module, just replace Module to your
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"Module":"game_8192"}]}]}

check activity of couple Modules
{"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"Any":[{"Module":"game_8192"}, {"Module":"devnet_nft"}]}]}

if you need any help to create your own filter feel free dming to @romanosadchyi Thanks!`

	Start = `
Hi, User! Just send json requsts via your filters like (currently on devnet network):
				
     {"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"MoveEvent"}, {"Package":"0x2"}, {"Module":"devnet_nft"}]}]}
     
     {"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[{"EventType":"Publish"}]}]}
     
     {"jsonrpc":"2.0", "id": 1, "method": "sui_subscribeEvent", "params": [{"All":[]}]}
	
Input /example to check more example.
More info are there https://docs.sui.io/build/event_api#event-filters`
}
