# go-backend-test

  O intuito desse desafio é avaliar seus conhecimentos técnicos em Golang e backend.

  Desenvolva duas aplicações em Golang: uma produtora e outra consumidora de mensagens.
  
  É obrigatório utilizar um sistema de filas para atualizar o status de transmissões de streaming em tempo real.

  Os estados das transmissões são: "Upcoming", "Live" e "Ended".

  O navegador não deve atualizar a página para obter o estado atual da transmissão mais próxima ou da que está acontecendo.
  
  Além disso, é necessário que tenha um endpoint para criar as transmissões e outro que consulte as que já aconteceram e retorne essas informações.

  Faça o fork desse repositório

## Requisitos
  - Escolha se os endpoints REST vão ficar no producer ou no consumer; 
  - Crie o endpoint de login e integre com o botão no frontend já criado.
  - Desenvolva o endpoint para criar a transmissão e o endpoint que busca as transmissões passadas, e integre com os botões do frontend já criado.
  - Crie o endpoint que traz o estado atual da transmissão mais próxima *no consumer* (Pelo horário e pela duração definida) e integre com o frontend já criado. 
  - Dê uma sugestão de arquitetura para fazer deploy dessa aplicação em um ambiente de produção e as razões pelas quais você escolheu ela;
  - Use container

## Envio
  Faça um Pull Request para esse repositório


## Prazo
  5 dias

## Avaliação

  - Sua aplicação preenche os requesitos solicitados?
  - Você documentou a maneira de configurar o ambiente e rodar sua aplicação?
  - Você seguiu as instruções de envio do desafio?
  - O seu código é organizado e legível?
  - Sua solução tem uma boa usabilidade?

## Diferenciais

  - Testes unitários e testes de integração;
  - Explicação das tecnologias utilizadas;
  - Utilizar boas práticas de segurança da informação;
  - Performance;