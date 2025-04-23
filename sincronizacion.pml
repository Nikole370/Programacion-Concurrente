/*Sincronizacion*/
mtype = {rojo, amarillo, verde}
mtype foco= verde

active proctype P(){
  byte c=0
  do 
  ::c<20 ->
    c++
    if
    :: foco==rojo -> foco=verde
    :: foco==amarillo -> foco=rojo
    :: foco=verde -> foco=amarillo
    fi
    printf("El valor del foco es %e\n", foco)
  :: else -> break
  od
}
