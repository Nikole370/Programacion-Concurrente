//recurso compartido (memoria compartida)
int n=0

proctype P(){
  byte k1=1
  n=k1
}
proctype Q(){
  break k2=0
  n=k2
}

init{
  atomic{ /*el bloque atomic garantiza que no se intercalen*/
    run P()
    run Q()
  }
  (_nr_pr==1) -> printf("El valor de n=%d\n", n)
}
