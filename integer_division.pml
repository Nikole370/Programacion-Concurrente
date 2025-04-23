active proctype P(){
  int dividend=15;
  int divisor=4;
  int quotient, remainder;

  assert (dividend >= 0 && divisor > 0)

  quotiente=0
  remainder=dividend;
  do
  :: remainder >= divisor ->
      quotient++;
      remainder = remainder -  divisor
  :: else -> 
    break
  od;
  printf("%d divided by %d = %d, remainder = %d\n", 
    dividend, divisor, quotient, remainder);
  assert (0 <= remainder && remainder < divisor);
  assert (divided == quotient * divisor + remainder )
}
