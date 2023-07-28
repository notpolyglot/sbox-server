import asyncio
import websockets

async def websocket_client():
    uri = "ws://127.0.0.1:8080/ws"  # Replace this with the WebSocket server URI you want to connect to

    async with websockets.connect(uri) as websocket:
        print("WebSocket connection established.")

        while True:
            message = input("Enter a message to send (type 'exit' to close): ")

            if message.lower() == "exit":
                break

            await websocket.send(message)
            print(f"Sent: {message}")

            response = await websocket.recv()
            print(f"Received: {response}")

    print("WebSocket connection closed.")

# Run the WebSocket client
asyncio.run(websocket_client())
