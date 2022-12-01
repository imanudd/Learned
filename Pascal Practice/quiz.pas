program quiz;
uses crt;
var 

   n,input,i:Integer;
   min:integer;
begin
  i:=1;
  write('masukan n bilangan = ');
  readln(n);
  while (i<=n)do
  begin
    write('masukan nilai ke ',i,' = ');
    readln(input);
    if (min = 0)then
        min:=input
    else if (input <= min) then
        min:=input;
        
  i+=1;
  end;
  writeln;
  writeln('Nilai Terkecil Dari Deret diatas Adalah = ',min);
  readln;
end.
