## Desenvolvimento de Aplicações Modernas e Escaláveis com Microsserviços

#### Kubernetes e hpa

1) Nessa fase, você deverá implementar exatamente o mesmo algorítimo que foi implementado em PHP (no php-apache-hpa / looping somando a raiz quadrada), porém na linguagem Go Lang. Não esqueça de seguir os passos abaixo:

    Nome do deployment, service e nome da imagem deverá se chamar: go-hpa
    Desenvolva os testes
    Implemente o processo de CI
    Faça o push da imagem no Docker Hub (seu-user/go-hpa)
    Faça o deploy da imagem no K8S (criando os arquivos de deployment e services)
    Cada réplica deverá consumir no mínimo 50m e no máximo 100m.

2) Implemente o "hpa" para que esse deployment tenha as seguintes características:

    O processo de escala inicia quando a CPU passar de 15%
    Quantidade mínima de pods: 1
    Quantidade máxima de pods: 6

3) Crie um POD e faça requisições através de um looping infinito e verifique se o autoscaler está funcionando corretamente.

  kubectl run -it  loader --image=busybox /bin/sh
    wget -q -O- http://go-hpa.default.svc.cluster.local;
    while true; do wget -q -O- http://go-hpa.default.svc.cluster.local; done;


Imagem:  
https://hub.docker.com/repository/docker/diovanes/go-hpa