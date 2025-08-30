# WAS Network Architecture

## 🏗️ Initial Entities

### Core Components
- **Broker** - Central communication hub
- **Agent** - Client nodes that connect to broker

---

## 🔒 Private Network Architecture

The WAS system operates within a private network environment, enabling secure communication between distributed components.

---

## 🔄 Initial Flow & Connection Process

### Phase 1: Broker Discovery & Agent Connection

#### Broker Initialization
1. **TCP Server Setup**
   - Broker opens TCP connection
   - Listens for incoming messages on designated port
   - Ready to accept agent connections

#### Agent Discovery Challenge
```
❓ Problem: How does an agent discover and connect to the broker?
```

#### Solution: UDP Broadcast Discovery
2. **Broker Advertisement**
   ```
   Protocol: UDP Broadcast
   Target: 255.255.255.255 (broadcast address)
   Message Format: "BROKER: [PORT]"
   Purpose: Announce broker presence to network
   ```

3. **Agent Listening**
   ```
   Agent listens on: 0.0.0.0 (all interfaces)
   Captures: Broker IP from broadcast packet
   Action: Extract broker IP and PORT information
   ```

4. **TCP Connection Establishment**
   ```
   Agent initiates: TCP connection to [brokerIP]:[PORT]
   Result: Persistent connection for communication
   ```

---

## 📁 Agent Connection Directory

### Connection Management System

#### Agent Registration Process
1. **Agent Identification**
   - Each message from agent includes unique `agentID`
   - Broker receives and processes agent identification

2. **Directory Storage**
   ```go
   // Conceptual data structure
   connectionMap := map[agentID]Peer{
       "agent_001": PeerConnection{...},
       "agent_002": PeerConnection{...},
       // ... more agents
   }
   ```

#### Benefits
- ✅ **Centralized Tracking** - Broker maintains complete agent registry
- ✅ **Message Routing** - Direct message delivery by agentID
- ✅ **Connection State** - Monitor agent connectivity status
- ✅ **Scalability** - Easy addition/removal of agents

---

## 🔌 Protocol Architecture

### Transport Layer Design

#### Current Implementation
```
Primary Transport: TCP
- Reliable delivery
- Connection-oriented
- Error handling built-in
```

#### Future Extensibility
```go
// Transport interface for future protocols
type Transport interface {
    Connect(address string) error
    Send(data []byte) error
    Receive() ([]byte, error)
    Close() error
}

// Implementations could include:
// - TCPTransport (current)
// - UDPTransport (future)
// - WSTransport (WebSocket)
// - QUICTransport (HTTP/3)
```

---

## 🔄 Complete Flow Diagram

```
1. Broker Startup
   ├── Open TCP listener on port X
   └── Start UDP broadcast: "BROKER: X"

2. Agent Discovery
   ├── Listen for UDP broadcasts (0.0.0.0)
   ├── Receive "BROKER: X" from [broker_ip]
   └── Extract broker IP and port

3. Connection Establishment
   ├── Agent → TCP connect([broker_ip]:X)
   ├── Send agentID in first message
   └── Broker stores in map[agentID]Peer

4. Operational State
   ├── Persistent TCP connections
   ├── Message routing by agentID
   └── Connection directory maintained
```


