# kafka-start

## Streaming de eventos

- Tudo é baseado em eventos
- Eventos podem ou devem ser armazenados e recuperados
- Toda comunicação baseada em eventos precisa ser escalável, resiliente e com alta disponibilidade

## Características

- Altíssimo throughput
- Latência extremamente baixa (2ms)
- Escalável e com alta disponibilidade
- Armazenamento
- Se conecta com quase tudo
- Construído em Java
- Bibliotecas para diversas tecnologias
- Open-source

## Quem usa

- Linkedin (nasceu lá ...)
- Bancos e/ou empresas de serviços financeiros
- Paypal, Netflix, Uber, Twitter, Dropbox, Spotify e outros

## Funcionamento

### Tópicos

- Canal que recebe e disponibiliza os dados)
- São parecidos com registros de logs com sequências e com ID (offset)
- Estes registros armazenados tem a seguinte anatomia (offset 0 = HEADERS / KEY / VALUE / TIMESTAMP)

### Partições

- Cada tópico pode ter uma ou mais partições para conseguir garantir a distribuição e resiliência de seus dados
- As partições permite que os registros estejam divididos, melhorando a escala do consumo deste registros
- A ordem das mensagens só serão garantidas se as mesmas estiverem na mesma partição
- As keys sevem para garantir a ordem, neste caso o kafka faz com que as mensagens com a mesma Key estejam na mesma partição
- No Kafka é possível replicar as partições em brokers diferentes através do replicator factor (Partições distribuídas)
- Com o replicator factor dizemos para o Kafka que não queremos perder as mensagens de um broker com problema (criticidade)
- Lembre-se que o replicator factor vai utilizar mais disco, memória, processamento e etc
- Para não ocorrer o problema de ler das duas partições o Kafka elege a partição líder
- As partições replicadas estarão disponíveis para leitura quando a partição líder estiver indisponível

### Produtores

- A garantia de entrega depende se queremos receber ou não a resposta de recebimento da mensagem pelo broker

#### Cenário Ack 0

- Se não queremos, usamos Ack 0 (acknowledge) para o comportamento de FF (fire-and-forget)
- Neste cenário não nos preocupamos com a garantia de recebimento pelo broker

#### Cenário Ack 1

- Neste modelo o broker envia a confirmação de recebimento para o produtor
- Mas se houver erro na replicação para os outros brokers é um falso positivo
- Este processo é mais lento pela fato da confirmação de recebimento pelo broker

#### Cenário Ack -1 (ALL)

- O produtor manda a mensagem o líder recebe e replica nos demais brokers
- O líder confirma para o produtor que a mensagem foi recebida
- Este é o cenário com a melhor garantia
- Este é o cenário mais lento

#### Alguns tipos de garantias

- At most once: garante a melhor performance mas pode perder algumas mensagens
- At least once: performance moderada, nao perde mensagems mas pode duplicar mensagens
- Exacly once: pior performance, nao perde mensagens e não duplica mensagens

#### Indepotência

- O Kafka consegue discartar uma mensagem duplicada, no modelo de indepotência ON
- Assim o Kafka coloca a mensagem na ordem correta (timestamp) e garente a entrega desta mensagem apenas 1x para o consumidor

### Consumidores

- Um consumidor pode ler as mensagens de uma ou de todas as partições de um tópico

#### Um grupo de consumidores

- Neste modelo o kafka organiza a entrega para os consumidores do mesmo grupo
- Se os consumidores não estao em um grupo, cada consumidor é o proprio grupo recebendo as mensagens de todas as partições
- Não tem como dois consumidores de um mesmo grupo fazerem a leitura de uma mesma partição
- Melhor que a quantidade de consumidores de um grupo seja a mesma quantidade das partições de um tópico

## Importante

- Sem uma KEY você não consegue ler as mensagens na ordem que as mesmas foram recebidas

## Comandos

- **Criando um tópico**
  - kafka-topics --create --topic=teste --bootstrap-server=localhost:9092 --partitions=3

- **Listando os tópicos**
  - kafka-topics --list --bootstrap-server=localhost:9092

- **Descrevendo um tópico**
  - kafka-topics --bootstrap-server=localhost:9092 --topic=teste --describe

- **Consumindo mensagens de um tópico**
  - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste

- **Publicando mensagens em um tópico**
  - kafka-console-producer --bootstrap-server=localhost:9092 --topic=teste

- **Consumindo mensagens desde o inicio**
  - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning

- **Consumindo mensagens através de grupo de consumidores**
  - kafka-console-consumer --bootstrap-server=localhost:9092 --topic=teste --from-beginning --group=g1

- **Descrevendo o grupo de consumidores**
  - kafka-consumer-groups --bootstrap-server=localhost:9092 --group=g1 --describe
