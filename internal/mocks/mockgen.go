package mocks

//go:generate sh -c "mockgen -package mockquic -destination quic/stream.go github.com/xyproto/quic Stream && goimports -w quic/stream.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/session.go github.com/xyproto/quic Session && goimports -w quic/session.go"
//go:generate sh -c "mockgen -package mockquic -destination quic/listener.go github.com/xyproto/quic Listener && goimports -w quic/listener.go"
//go:generate sh -c "../mockgen_internal.sh mocks short_header_sealer.go github.com/xyproto/quic/internal/handshake ShortHeaderSealer"
//go:generate sh -c "../mockgen_internal.sh mocks short_header_opener.go github.com/xyproto/quic/internal/handshake ShortHeaderOpener"
//go:generate sh -c "../mockgen_internal.sh mocks long_header_opener.go github.com/xyproto/quic/internal/handshake LongHeaderOpener"
//go:generate sh -c "../mockgen_internal.sh mocks crypto_setup.go github.com/xyproto/quic/internal/handshake CryptoSetup"
//go:generate sh -c "../mockgen_internal.sh mocks stream_flow_controller.go github.com/xyproto/quic/internal/flowcontrol StreamFlowController"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/sent_packet_handler.go github.com/xyproto/quic/internal/ackhandler SentPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mockackhandler ackhandler/received_packet_handler.go github.com/xyproto/quic/internal/ackhandler ReceivedPacketHandler"
//go:generate sh -c "../mockgen_internal.sh mocks congestion.go github.com/xyproto/quic/internal/congestion SendAlgorithmWithDebugInfos"
//go:generate sh -c "../mockgen_internal.sh mocks connection_flow_controller.go github.com/xyproto/quic/internal/flowcontrol ConnectionFlowController"
