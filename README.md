# UpdateCompose


## Install

```bash
bash <(curl -Ls https://owo.nz/update-compose)
```


## Run

```bash
update-compose
```

output:
```
Found Compose file in directory /root/azure2openai
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/azure2openai
Found Compose file in directory /root/chatgpt
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/chatgpt
Found Compose file in directory /root/cohere2openai
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/cohere2openai
Found Compose file in directory /root/deeplx
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/deeplx
Found Compose file in directory /root/discord-image
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/discord-image
Found Compose file in directory /root/freeduckduckgo
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/freeduckduckgo
Found Compose file in directory /root/lobe-chat
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/lobe-chat
Found Compose file in directory /root/one-api
Running command: docker compose stop
Running command: docker compose pull
Running command: docker compose up -d
Updated compose in /root/one-api
```