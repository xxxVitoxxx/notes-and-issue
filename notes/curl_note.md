# curl

## curl 是什麼?

curl 是一個開源專案，基於網路協定，對指定 `URL` 進行網路傳輸，不涉及資料的處理  

## 常見用法

- 回傳 `URL` 的 response ，根據 `URL` 內容，可能是 HTML 、 JSON 或 XML 格式  

```linux
curl https://www.google.com  
```  

- `-o`  
  搭配自定義的檔名及網址，下載檔案  

```linux
curl -o picture.png  https://cdn-icons-png.flaticon.com/512/3067/3067260.png
```  

- `-O`  
  搭配網址，會直接下載並用原始的檔名儲存  

```linux
curl -O https://cdn-icons-png.flaticon.com/512/3067/dog1.png
```  

- 迴圈下載  
  將檔名改成 `dog[1-5].png`，就會依序下載 `dog1.png` 、 `dog2.png` 到 `dog5.png`  

- `-x`  
  透過 proxy 及 port 使用 curl  

```linux
curl -x {proxy_host}:{proxy_port} https://www.google.com
```  

- `-c`  
  儲存 response 的 `cookie` 資訊  

```linux
curl https://www.google.com -c cookie.txt
```  

- `-D`  
  儲存 response 的 `header` 資訊  

```linux
curl https://www.google.com -D header.txt
```

- `-b`  
  使用儲存的 `cookie` 訪問網站，很多網站都是通過監視你的 `cookie` 資訊判斷是否按規矩訪問網站  

```linux
curl -b {cookie_file} https://www.google.com
```  

- 模仿瀏覽器  
  有些網站需要使用特定的瀏覽器或特定版本才能成功訪問，可以使用 `-A` 指定瀏覽器  
  以下範例會讓 server 端認為你是用 `IE8.0` 訪問的  

```linux
curl -A "Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 5.0)" https://www.google.com
```  

- `-H`  
  將 `header` 資訊傳給 server 端  

```linux
curl -H "Content-Type:application/json" https://www.google.com
```  

- `-I` 、 `--head`  
  只顯示文件資訊  

- `-d`  
  `POST` 的資料  

```linux
curl -d "{"message":"hello, world"}" https://www.google.com
```  

- `-u`  
  server 端的帳密  

```linux
curl -u {account}:{password} https://www.google.com
```  
