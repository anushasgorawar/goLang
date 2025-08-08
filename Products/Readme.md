1.
http.HandleFunc("/product/id", returnOneProduct)
func returnOneProduct(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
}
```
2025/08/08 02:07:46 &{GET /product/id HTTP/1.1 1 1 map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7] Accept-Encoding:[gzip, deflate, br, zstd] Accept-Language:[en-US,en;q=0.9,kn-IN;q=0.8,kn;q=0.7,hi-IN;q=0.6,hi;q=0.5] Connection:[keep-alive] Sec-Ch-Ua:["Not)A;Brand";v="8", "Chromium";v="138", "Google Chrome";v="138"] Sec-Ch-Ua-Mobile:[?0] Sec-Ch-Ua-Platform:["macOS"] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36]] {} <nil> 0 [] false localhost:9090 map[] map[] <nil> map[] 127.0.0.1:63079 /product/id <nil> <nil> <nil> /product/id 0xc000132320 0xc00006a360 [] map[]}
```

2. http.HandleFunc("/product/", returnOneProduct)
```
2025/08/08 02:10:37 &{GET /product/id HTTP/1.1 1 1 map[Accept:[text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7] Accept-Encoding:[gzip, deflate, br, zstd] Accept-Language:[en-US,en;q=0.9,kn-IN;q=0.8,kn;q=0.7,hi-IN;q=0.6,hi;q=0.5] Connection:[keep-alive] Sec-Ch-Ua:["Not)A;Brand";v="8", "Chromium";v="138", "Google Chrome";v="138"] Sec-Ch-Ua-Mobile:[?0] Sec-Ch-Ua-Platform:["macOS"] Sec-Fetch-Dest:[document] Sec-Fetch-Mode:[navigate] Sec-Fetch-Site:[none] Sec-Fetch-User:[?1] Upgrade-Insecure-Requests:[1] User-Agent:[Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/138.0.0.0 Safari/537.36]] {} <nil> 0 [] false localhost:9090 map[] map[] <nil> map[] 127.0.0.1:63091 /product/id <nil> <nil> <nil> /product/ 0xc0001040a0 0xc0000ae2a0 [] map[]}
2025/08/08 02:10:37 Endpoint hit: Homepage
```

log.Println(r.URL.Path)
```
2025/08/08 02:11:58 /product/1
```


Problem with this is path: path := r.URL.Path[len("/product/"):]

Better way: Gorilla mux
1. Supports Method Based Routing
- Router makes it easy to dispatch a http request to handlers based on the methods such as GET, PUT, POST, DELETE and PATCH
2. Supports Variables in URL paths
- eg. /movies/{id} where id is a dynamic variable
