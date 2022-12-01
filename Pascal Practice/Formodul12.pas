program Formodul12;
uses crt;

function soal1():integer;
var i :Integer;
begin
  writeln('Bilangan Kelipatan 2, 0-100');
  for i:=1 to 100 do
  begin
    if(i mod 2 = 0)then
    write(i,' ');
  end;
end;

function soal2():integer;
var i,a:integer;
begin
  write('buat angka batasan = ');readln(i);
  for a:=1 to i do
  begin
    write(a,' ');
  end;

end;

begin
  clrscr;
  writeln('***********************************');
  soal1();
  writeln;
  writeln('***********************************');
  soal2();
  writeln;
  writeln('***********************************');
  readln;
end.
