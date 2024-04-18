# Creación de monitores - IX

### Contexto:

En 2023, el equipo de IX tenía un squad que se encargaba de la creación y tuneo de monitores transaccionales para aquellas marcas pertenecientes al Tier definido por el equipo comercial. Estos monitores tienen como finalidad alertar al NOC ante una eventual caída de las marcas para que este realice el análisis respectivo y así detectar si se debe a un problema interno de Mercado Pago, o si de lo contrario se debe informar a la marca para que revisen si existe alguna afectación de su lado.

Debido a cambios de estructura en IX, la estructura de los squads finalizó y se precisa desarrollar una funcionalidad que permita automatizar la creación de los monitores, hacer más escalable su creación de los monitores a demanda del equipo comercial, sin precisar de una intervención manual por parte de algún integrante de IX.

### Objetivo: 

Crear una herramienta / automatización que se encargue de crear de manera automática los monitores de marcas que no sean TAM, bajo las siguientes condiciones:
Los monitores a crear serán definidos y solicitados por el equipo comercial y/o el equipo de integraciones (IX).
La creación de los monitores se realizará una vez al mes en una fecha específica (aún por definir), en la cual se ejecutará el proceso con las solicitudes obtenidas antes de esta fecha.
Los monitores serán creados en DataDog con una configuración tipo threshold para facilitar el tuning.

### Requerimientos funcionales:

La herramienta deberá contar con las siguientes características:
- Un formulario de entrada donde los comerciales puedan llenar la información necesaria para la creación del monitor. Esta información es:
    - Business Unit (obligatorio)
    - Producto (obligatorio)
    - Nombre de la marca (obligatorio)
    - Flujo (opcional): Descripción del flujo transaccional de la marca. Ej: B2C
    - Site (obligatorio)
    - Horario de atención al cliente de la marca (obligatorio)
    - ¿Plataforma/PDV? (obligatorio): Será un booleano, en el cual se debe especificar si lo que se quiere monitorear es una plataforma/PDV. Si se marca false indicará que lo que se quiere monitorear es una marca.
    - Identificadores de la marca (obligatorio). Los campos a escoger son:
        - Cust IDs (Opción disponible obligatoria solo para OP - a excepción de wallet connect WC - si Plataforma/PDV false)
        - App ID / Client ID (Opción disponible opcional para todos los productos de OP menos WC y obligatoria para Delivery y WC si Plataforma/PDV false)
        - Marketplace ID (Opción disponible si Plataforma/PDV true)
        - Platform ID (Opción disponible si Plataforma/PDV true)
        - Brand ID (Opción disponible solo para QR y Point si Plataforma/PDV false)
        - Sponsor ID (Opción disponible si Plataforma/PDV true)
    - TPV Mensual en USD (obligatorio)
    - KAM (obligatorio)
    - Contacto Seller (opcional?)
    - Contacto Técnico Seller (obligatorio)

- Se creará un cronjob que se ejecutará una vez al mes (configurable) y consumirá un servicio API Rest (microservicio), el cual se encargará de crear los monitores solicitados si estos no existen. La creación de estos monitores se hará usando las métricas correspondientes al agrupador y producto de la siguiente manera:
    - Plataforma/PDV true:
    - Marketplace ID: integrations.marketplace.payments
    - Platform ID: integrations.platforms.payments
 - Sponsor ID: integrations.platforms.payments
Plataforma/PDV false:
Cust IDs: integrations.payments con filtro para especificar la unidad de negocio que es business_unit:online_payments y filtro de application_id si este fue completado. Debido a que este criterio será solo necesario para OP (exceptuando WC). El filtro que se usará para esta métrica será el filtro de business_group.
App ID / Client ID: 
Si es Wallet Connect (WC) el filtro a utilizar será client_id y se deberán crear dos monitores, uno de advanced payments y otro de confirmación de agreements. Cada uno con su métrica correspondiente:  advanced_payments.ap_counter y connecting_tools.wallet_agreement_manager.agreement_confirmed
Si es Delivery el filtro a utilizar será application_id: proximity_integration.shipment.request.metric
Brand ID:
Si es QR: integrations.qr.payments
Si es Point: integrations.point.payments
Estos además tendrán configurados los downtime correspondientes según los horarios de servicio que se obtuvieron en el formulario.	
El microservicio que ejecuta el cronjob, realizará un conjunto de validaciones previas a la creación de los monitores para evitar la duplicación de los mismos. Se ejecutará entonces el siguiente flujo de validación:
Para Plataforma/PDV false:
Para OP exceptuando Wallet Connect (WC):
Validar existencia de business_group: Deberá validarse si los collectors pertenecen a un BG. Si no existe deberá crearlo.
Si el BG existía previamente, validar la existencia de un monitor: Se deberá validar si ya existe un monitor para el identificador o conjunto de agrupadores escogidos (ya sea BG y application_id si fue completado, o solo BG si el application_id no fue llenado).
Para QR y Point:
Validar la existencia del monitor: Se deberá validar si ya existe un monitor para el brand_id indicado para las métricas de QR y Point según corresponda.
Si el monitor no existe, validar la existencia del brand_id en la base de datos de brands y en el API de metric pusher para QR o Point según corresponda. Si el brand_id no existe en la base de datos de brands, el flujo termina y se notifica al solicitante de dicho error. Si el brand_id existe en la base de datos de brands pero no en el API de metric pusher, agregar la brand al API de metric pusher del producto correspondiente.
Para Wallet Connect y Delivery:
Validar la existencia del monitor.
Para Plataforma/PDV true:
Validar la existencia del monitor (Independiente del producto): Se deberá validar la existencia del monitor según la métrica propia del agrupador.
Validar si el agrupador está habilitado para que se le pushee data:
Para los casos de los agrupadores platform_id y sponsor_id: Validar si el agrupador está disponible en el ADP. Si no existe, el flujo termina y se notifica al solicitante de dicho error.
Para los casos del agrupador marketplace_id: Validar si el agrupador está disponible en el API de metric pusher de marketplace. Si no existe, agregarlo.
Una vez realizadas todas las validaciones correspondientes, si los monitores no existen y cumple de manera exitosa las validaciones descritas, procederá a crear los monitores. Así mismo deberá validar la existencia de los downtimes, y si no existen crearlos para configurarlos a cada monitor.
El microservicio y el cronjob  deben contar con y guardar logs (Datadog, NewRelic, Kibana).
Una vez termina el flujo de intento de creación del monitor del script, se genera una notificación del resultado de la creación. Esta notificación tendrá las siguientes características:
Información que se usó para la creación del monitor (es decir la que se llenó en el formulario).
En caso de éxito:
Se enviará al solicitante 
Link del monitor creado
Fecha de creación
En caso de error:
Se enviará al solicitante
Descripción del error obtenido
Pasos para corregir el error. Ej: Si el flujo falló porque no existe el brand_id solicitado, pedirle al solicitante que se asegure de que el brand_id exista en el admin de brands.
Cada solicitud debe crear un registro en una base de datos, en la cual se consignará:
ID de solicitud
Fecha de creación de la solicitud
Información proporcionada en la solicitud
ID del monitor creado (si fue exitoso)
Error (si falló el proceso)
Fecha de creación del monitor (o fecha de error según el caso)
Envío exitoso de notificación (boolean)
Error del envío de la notificación si aplica
Receptores de notificación
Fecha de envío de la notificación (o de error)

Nota: Se espera que las notificaciones de la creación de monitores se realicen a un canal de Slack en donde se pueda además arrobar al solicitante. 











Criação de monitores - IX
Contexto
Em 2023, a equipe IX contava com uma squad encarregada de criar e ajustar monitores transacionais para as marcas pertencentes ao Tier definido pela equipe comercial. Esses monitores têm como objetivo alertar o NOC caso haja uma possível queda nas marcas para que ele possa realizar a respectiva análise e assim detectar se é devido a um problema interno do Mercado Pago, ou se de outra forma a marca deve ser informada para que eles podem revisá-lo se houver alguma afetação da sua parte.

Devido a mudanças estruturais no IX, a estrutura do squad terminou e é necessário desenvolver uma funcionalidade que permita automatizar a criação de monitores, tornando a criação de monitores mais escalável a pedido da equipe de vendas, sem exigir intervenção manual por parte de membro do IX.
Objetivo 
Construir uma ferramenta/automação responsável por criar automaticamente monitores para outras marcas que não a TAM, nas seguintes condições:
Os monitores a serem criados serão definidos e solicitados pelo time comercial e/ou pelo time de integrações (IX).
A criação dos monitores será realizada uma vez por mês em data específica (ainda a ser definida), na qual será executado o processo com as solicitações obtidas antes desta data.
Os monitores serão criados no DataDog com configuração do tipo limite para facilitar o ajuste.


Requerimentos funcionais:

A ferramenta deve ter as seguintes características:
Formulário de inscrição onde os vendedores podem preencher as informações necessárias para a criação do monitor. Esta informação é:
Business Unit (obrigatório)
Producto (obrigatório)
Nombre de la marca (obrigatório)
Flujo (opcional): Descripción del flujo transaccional de la marca. Ej: B2C
Site (obrigatório)
Horario de atención al cliente de la marca (obrigatório)
¿Plataforma/PDV? (obrigatório): Será um booleano, no qual você deverá especificar se o que deseja monitorar é uma plataforma/POS. Se false estiver marcado, indicará que o que você deseja monitorar é uma marca.
Identificadores de la marca (obrigatório). Os campos a escolher são:
Cust IDs (Opção obrigatória disponível apenas para OP - a excepción de wallet connect WC - si Plataforma/PDV false)
App ID / Client ID (Opção disponível opcional para todos os produtos OP exceto WC e obrigatória para Delivery e WC se Plataforma/POS falso)
Marketplace ID (Opção disponível se Plataforma/PDV true)
Platform ID (Opção disponível se Plataforma/PDV true)
Brand ID (Opção disponível só para QR e Point se Plataforma/PDV false)
Sponsor ID (Opção disponível sePlataforma/PDV true)
TPV Mensual en USD (obrigatório)
KAM (obrigatório)
Contacto Seller (opcional?)
Contacto Técnico Seller (obrigatório)
Será criado um cronjob que será executado uma vez por mês (configurável) e consumirá um serviço Rest API (microsserviço), que será responsável por criar os monitores solicitados caso eles não existam. A criação destes monitores será feita utilizando as métricas correspondentes à garoupa e produto da seguinte forma:
Plataforma/PDV true:
Marketplace ID: integrations.marketplace.payments
Platform ID: integrations.platforms.payments
Sponsor ID: integrations.platforms.payments
Plataforma/PDV false:
Cust IDs: integrations.payments com filtro para especificar a unidade de negócios que está business_unit:online_payments e filtro de application_id se isso foi concluído. Porque este critério só será necessário para OP (exceto WC). O filtro que será utilizado para esta métrica será o filtro business_group.
App ID / Client ID: 
Se for Wallet Connect (WC), o filtro a utilizar será client_id e deverão ser criados dois monitores, um para adiantamentos e outro para confirmação de acordos. Cada um com sua métrica correspondente: advanced_payments.ap_counter y connecting_tools.wallet_agreement_manager.agreement_confirmed
Se for Delivery, o filtro a utilizar será application_id: proximity_integration.shipment.request.metric
Brand ID:
Se é QR: integrations.qr.payments
Se é Point: integrations.point.payments
Estes também terão o tempo de inatividade correspondente configurado de acordo com os horários de atendimento obtidos no formulário.	
O microsserviço que executa o cronjob realizará um conjunto de validações antes da criação dos monitores para evitar sua duplicação. O seguinte fluxo de validação será então executado:
Para Plataforma/PDV false:
Para OP exceto Wallet Connect (WC):
Validar a existência de business_group: Deve ser validado se os coletores pertencem a um BG. Se não existir, você deverá criá-lo.
Se o BG existia anteriormente, validar a existência de um monitor: Deve ser validado se já existe um monitor para o identificador ou conjunto de garoupas escolhido (ou BG e application_id se foi preenchido, ou apenas BG se o application_id não foi preenchido). 
Para QR y Point:
Validar a existência do monitor: Deve ser validado se já existe um monitor para o brand_id indicado para métricas QR e Point conforme apropriado.
Caso o monitor não exista, valide a existência do brand_id no banco de dados de marcas e na API do metric pusher para QR ou Point conforme apropriado. Caso o brand_id não exista no banco de dados de marcas, o fluxo termina e o solicitante é notificado deste erro. Se brand_id existir no banco de dados de marcas, mas não na API do metric pusher, adicione a marca à API do metric pusher do produto correspondente.
Para Wallet Connect y Delivery:
Valide a existência do monitor.
Para Plataforma/PDV true:
Validar a existência do monitor (Independente do produto): A existência do monitor deve ser validada de acordo com a métrica própria do agrupador.
Valide se o costurador está habilitado para dados do pushee:
Para os casos dos garoupas platform_id e patrocinador_id: Validar se o agrupador está disponível no ADP. Se não existir, o fluxo termina e o solicitante é notificado desse erro.
Para os casos do garoupa marketplace_id: Valide se o agrupador está disponível na API do marketplace metric pusher. Se não existir, adicione-o.
Una vez realizadas todas las validaciones correspondientes, si los monitores no existen y cumple de manera exitosa las validaciones descritas, procederá a crear los monitores.Da mesma forma, deve-se validar a existência dos tempos de inatividade e, caso não existam, criá-los para configurá-los para cada monitor.
O microsserviço e o cronjob devem possuir e salvar logs (Datadog, NewRelic, Kibana).
Depois que o fluxo de tentativa de criação do monitor de script for concluído, uma notificação do resultado da criação será gerada. Esta notificação terá as seguintes características:
Informações que foram utilizadas para criar o monitor (ou seja, o que foi preenchido no formulário).
Em caso de successo:
Será enviado ao requerente: 
Link do monitor criado;
Data de criação
Em caso de erro:
Será enviado ao requerente
Descriçã do erro 
Passos para corrigir o erro. Ex: Se o fluxo falhar porque o brand_id solicitado não existe, peça ao solicitante para garantir que o brand_id exista no administrador de marcas.
Cada solicitação deverá criar um registro em um banco de dados, que registrará:
ID de solicitud
Data de criação da solicitud
Informação fornecida na solicitud
ID do monitor criado (si foi um successo)
Erro (se falhou no processo)
Data de criação do monitor (ou data de erro conforme o caso)
Envio de notificação com sucesso (boolean)
Erro no envio de notificação, se aplicável
Receptores de notificação
Data de envio de notificación (ou erro)

Observação: Espera-se que as notificações da criação de monitores sejam feitas para um canal do Slack onde o solicitante também possa ser endossado.

