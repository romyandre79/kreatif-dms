import asyncio
import logging
from datetime import datetime
from fastapi import WebSocket

class ConnectionManager:
    def __init__(self):
        self.active_connections: list[WebSocket] = []

    async def connect(self, websocket: WebSocket):
        await websocket.accept()
        self.active_connections.append(websocket)

    def disconnect(self, websocket: WebSocket):
        if websocket in self.active_connections:
            self.active_connections.remove(websocket)

    async def broadcast(self, message: str):
        for connection in self.active_connections:
            try:
                await connection.send_text(message)
            except:
                pass

manager = ConnectionManager()
main_loop = None

async def broadcast_log(message: str, level: str = "info"):
    timestamp = datetime.now().strftime("%H:%M:%S")
    formatted_msg = f"{timestamp} [{level.upper()}] ocr-service: {message}"
    await manager.broadcast(formatted_msg)

# Custom Log Handler
class WebSocketLogHandler(logging.Handler):
    def emit(self, record):
        try:
            log_entry = self.format(record)
            if main_loop and main_loop.is_running():
                main_loop.call_soon_threadsafe(
                    lambda: asyncio.create_task(manager.broadcast(log_entry))
                )
        except Exception:
            self.handleError(record)
