## Teste de Código: MiniPay

### Objetivo

O **MiniPay** é uma plataforma de pagamentos simplificada. Nela, é possível realizar depósitos e transferências de dinheiro entre usuários. Existem **dois tipos de usuários**:

- **Comuns**: Podem transferir e receber dinheiro.
- **Lojistas**: Somente recebem dinheiro, não enviam.

### Requisitos

Abaixo estão as regras de negócio que são importantes para o funcionamento do PicPay Simplificado. Tente implementar as funcionalidades de forma aderente ao que é solicitado, mas não se preocupe se não conseguir atender a todos os requisitos. Durante a entrevista, vamos discutir o que foi possível fazer e o que não foi.

1. **Cadastro de Usuários**:
   - **Campos obrigatórios**: Nome Completo, CPF, E-mail, Senha.
   - **Validação**: CPF/CNPJ e e-mails devem ser únicos no sistema. Portanto, deve ser permitido apenas um cadastro com o mesmo CPF ou e-mail.

2. **Transferências**:
   - Usuários podem transferir dinheiro para lojistas e entre si.
   - **Lojistas** só podem receber transferências (não podem enviar dinheiro).
   - **Validação de saldo**: O sistema deve validar se o usuário tem saldo suficiente antes de permitir a transferência.
   
3. **Serviço Autorizador Externo**:
   - Antes de finalizar a transferência, consulte um serviço autorizador externo. Use o seguinte mock para simular a consulta:  
     [Mock do Autorizador Externo](https://run.mocky.io/v3/5794d450-d2e2-4412-8131-73d0293ac1cc)

4. **Transação Revertida em Caso de Erro**:
   - A operação de transferência deve ser realizada em uma **transação**. Caso haja algum erro, a operação deve ser revertida e o dinheiro deve voltar para a carteira do usuário pagador.

5. **Notificação de Recebimento de Pagamento**:
   - Após o recebimento de pagamento, o usuário ou lojista precisa ser notificado (via e-mail, SMS, etc.).
   - O serviço de notificação pode estar **indisponível ou instável**. Use o seguinte mock para simular o envio de notificações:  
     [Mock de Serviço de Notificação](https://run.mocky.io/v3/54dc2cf1-3add-45b5-b5a9-6bf7e7f1f4a6)

6. **API RESTful**:
   - O sistema deve expor uma **API RESTful** para que as operações de transferência possam ser realizadas.

---