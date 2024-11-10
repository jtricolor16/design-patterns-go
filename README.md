# Design Patterns - Golang

## Objetivo

Trata-se de um repositório criado para fins educativos. O propósito foi aplicar o conteúdo estudado em [design-patterns](https://refactoring.guru/design-patterns).

## Design Patterns

Há 3 grandes grupos nos quais se dividem os padrões de projeto.

- **Criacionais**: Lidam com o fluxo de criação de instâncias de objetos. Neste repositório, estão contemplados
   - **Singleton**: Responsável por garantir que somente uma única instância da classe irá ser definida durante todo o fluxo de execução;
   - **Factory Method**: Como o próprio nome sugere, preocupa-se em definir uma classe que terá como responsabilidade ser uma instanciadora de objetos, de acordo com a necessidade do invocador.
   - **Abstract Factory Method**: Trata-se de um refinamento aplicado ao *Factory Method* com o propósito de permitir o agrupamento de componentes que devem ser criados em conjunto. Para tal, mais uma interface é acrescida ao processo, com o intuito de simplificar este agrupamento e de também tornar tal comportamento extensível;
   - **Prototype**: Permite efetuar o clone de um objeto;
   - **Builder**: Através de uma estrutura orquestradora, efetua a criação de objetos necessários à execução de um contexto. Tais instâncias são provisionadas sequencialmente.

- **Estruturais**: Cuidam da integração entre diferentes objetos:
  - **Façade**: Inclui uma abstração para que, por exemplo, uma biblioteca não seja utilizada diretamente. Assim, garantindo um maior controle à aplicação consumidora da lib (principalmente quando a versão mudada é do tipo minor).
  - **Proxy**: Similar ao façade. Possibilita o acréscimo de recursos ou simplesmente ajustes ao retorno de uma biblioteca para que se adeque ao que a aplicação necessita.
  - **Decorator**: Inclusão dinâmica de comportamentos a uma determinada funcionalidade.
  - **Composite**: Uso de composição para estender o papel de uma classe.
  - **Adapter**: Cria um adaptador ao retorno de uma determinada requisição para que se adeque ao propósito da classe em questão.
  - **Bridge**: Através da relação entre abstração (interface de usuário, por exemplo) e implementações (componentes visuais que irão compor a interface visual), possibilita a integração entre os componentes mencionados com o propósito de enriquecer o objeto com a menor repetição de código possível.
  - **Flyweight**: Seu propósito principal é economizar a memória utilizada na aplicação. Contudo, para tal, a complexidade de código aumenta significativamente. Para alcançar isso, evita-se gerar diferentes instâncias de um mesmo objeto sem que haja necessidade.

- **Comportamental**: Por fim, os padrões comportamentais representam o modo como objetos distintos impactam no comportamento de outros objetos. Ou seja, podem efetuar alguma transição em tempo de execução ou, simplesmente, via composição, executar alguma funcionalidade atrelada a outra instância. Vamos aos padrões:
  - **Strategy**: Em tempo de execução, transiciona o objeto que será executado. Para tal, as instâncias precisam ter a mesma assinatura.
  - **Observer**: Similar ao comportamento do padrão pubsub. Ou seja, o objeto fica ouvindo até que seja sensibilizada a sua execução pela classe de subscrição.
  - **Mediator**: Centraliza a gestão dos steps a serem sensibilizados em uma classe concreta (mediadora) e estende a implementação das notificações nos componentes concretos que serão mediados pela classe anteriormente mencionada (cada uma dessas classes irá registrar o mediador).
  - **Memento**: Padrão utilizado quando queremos garantir o histórico de execuções para que possamos ativar o fluxo de retrocesso de fluxo.
  - **Chain of Responsability**: Encadeamento de execuções com o propósito de sincronizar os processos via uma mesma interface.
  - **State**: De maneira orquestrada, faz com que um objeto, via composição, tenha ciência do outro e efetue a troca de estado (comportamento similar ao de uma máquina de estado - orquestração)
  - **Visitor**: Padrão implementado com o propósito de garantir que um comportamento possa ser estendido a diferentes classes sem que o princípio de responsabilidade única do objeto da camada negocial seja violado.
  - **Template-Method**: Via sobrecarga de método, herança ou interface: proporciona a redução de código escrito quando se almeja efetuar pequenos ajustes entre objetos com certa similaridade comportamental.
  - **Command**: Garantindo um fluxo unidirecional, cria uma camada intermediária entre o cliente e a funcionalidade a ser executada com o propósito de reduzir o acoplamento entre ambos.git remote add origin git@github.com:jtricolor16/design-pattern-go.git