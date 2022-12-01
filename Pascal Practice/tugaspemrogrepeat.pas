program tugaspemrogrepeat;
uses crt;
var
jumlah,data:Integer;
j_data:integer;

function metode1():integer;
begin
repeat
  writeln('masukan data = ');
  readln(data);
  jumlah := jumlah+data;
  j_data+=1;
  until jumlah >= 50;
  writeln('---------------------------------');
  Write('Banyak nya data = ',j_data);
  write(' Jumlahnya ',jumlah);
  readln;
end;
begin
  clrscr;
  metode1();
end.
