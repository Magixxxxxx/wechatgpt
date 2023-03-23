import requests

token = "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"
headers = {
     "Authorization": "Bearer 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",
     "Content-Type": "application/json"
   }
proxies = {
    'https': 'http://sz.oneconnect.shopeemobile.com:6443', 
}
data = '{\
     "model": "gpt-3.5-turbo",\
     "messages": [{"role": "user", "content": "蹭蹭猫猫"}],\
     "temperature": 0.7\
   }'
resp = requests.post(url="https://api.openai.com/v1/chat/completions", headers=headers, data=data.encode("UTF-8"))
print(resp)