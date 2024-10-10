package service

type Service struct {
	KafkaProducer *KafkaProducer
	// kafkaConsumer *KafkaConsumer
}

func NewService(brokers, topic string) *Service {
	return &Service{
		KafkaProducer: NewKafkaProducer(brokers, topic),
		// kafkaConsumer: NewKafkaConsumer(brokers, groupID),
	}
}
