# Cli de gerenciamento servidor

> Cli desenvovida para ser consumida pela [Direcao Marcas E Patentes](https://direcaomarcas.com.br/), por isso algumas questões estão executando de forma não convencional

## Objetivos

- Facilitar o processo de Deploy e Execuções de funções no servidor VPS utilizado pela empresa
- Criar melhor conhecimento da linguagem de programação Go

### Requisitos Funcionais

- Fazer a configuração do git com a chave SSH
    - Atualmente é necessário toda vez que precisamos fazer alguma coisa em algum repositorio no git é necessário configurar a chave ssh que será utilizada. A ideia é que está CLI resolva este problema em qualquer repositório que ela seja executada
      - [ ] Executar os comandos `eval "$(ssh-agent -s)"` e `ssh-add ~/.ssh/github` no repositorio atual para que seja executada a configuração da chave
  
- Restart e rebuild dos containers
  - Atualmente devida a uma má configuração dos containers é necessário excluir antes de fazer a build para o deploy essa flag vai resolver isso

- Monitoramento de servicos
  - Vai monitorar os serviços e containers que precisam estar rodando, e vai notificar de alguma forma
    - [ ] Criar um log de monitoramento
    - [ ] Criar uma flag de adicionar servicos
    - [ ] Criar flag de adicionar container
    - [ ] Criar forma de comunicar o time sobre o não funcionamento de algum container
    - [ ] Fazer utilizando Crontab