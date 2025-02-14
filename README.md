# Dog-balancer

Esta aplicação se trata de uma versão simplificada de um load balancer, usando a estratégia **round-robin**.

## Algoritmo

Round-robin é considerada a estratégia mais simples quando se trata de balanceamento de recursos. Ela consiste em dividir de forma igualitária a quantidade de processamento entre threads, cpu's ou, no nosso caso, serviços. Primeiro, um `slice` é criado no arquivo `main.go`, contendo os endereços das instâncias de backend:

```go
var (
	backends = []string{
		"http://localhost:8001",
		"http://localhost:8002",
		"http://localhost:8003",
		"http://localhost:8004",
	}

	currentIndex = 0
)
```

Há também a variável currentIndex, armazenando um inteiro que se refere a instância de backend atual no round-robin.

Para seguir o algoritmo do round-robin, devemos entregar sempre uma requisição para cada serviço, fazendo com que todos recebam o mesmo número. É a única coisa que importa, sem levar em consideração estratégias de prioridade ou outras questões mais complexas:

```go
func IncrementSliceIndex(index *int, slice []string) {
	lastIndex := len(slice) - 1

	if *index == lastIndex {
		*index = 0

		return
	}

	*index++
}
```

A função `IncrementSliceIndex` trás o comportamento necessário para alternar de forma consistente o índice do `slice` recebido.

## Evitando race conditions

Sempre que uma requisição é feita, uma nova `goroutine` é agendada. Isso trás consigo a preocupação de compartilhamento de recursos. Para evitar um comportamento inesperado, proveniente de uma condição de corrida, foi utilizado um `mutex`, garantindo a leitura e escrita de forma segura do index atual:

```go
mu.Lock()

utils.IncrementSliceIndex(currentIndex, slice)

mu.Unlock()
```

## Futuro

Até então, o `dog-balancer` não tem suporte a configurações de TLS, então só pode ser usado com http. Adicionar esse suporte pode ser interessante. **Devo dizer que este é um projeto voltado para estudos, dog-balancer NÃO deve ser usado em produção.** Nunca nem se quer deve sonhar em substituir um nginx da vida. PR's são bem vindos.
