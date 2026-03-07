## How to Install

### 1. Build the Docker Image

```bash
docker build -t bot-tele-bps:1 .
```

### 2. Run the Container (Directly)

Run the container and start the bot immediately on background:

```bash
docker run -d --name bot-tele-bps bot-tele-bps:1
```

### 3. Run and Access the Container Shell

Run the container in interactive mode to access the shell:

```bash
docker run -it bot-tele-bps:1 sh
```

Then start the bot manually:

```bash
./bot-tele
```
