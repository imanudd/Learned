program algoritma;
uses Crt;
var 
    n:Integer;
    faktorial:real;
    index:Integer;
    
    x,y:Integer;
    pangkat:Real;

    a,r,n1:real;
    hasil: real;
    
    batas :real;
    first,sec:real;
    hasil2:real;
    kpk1,kpk2:real;

function faktorialn():real;
begin
  Write('Masukan Nilai N! = ');
  ReadLn(n);
  index:=1;
  faktorial:=1;
  while (index<=n) do
    begin
    faktorial:=faktorial*index;
    index+=1;
    end;
   WriteLn(faktorial);
end;

function xpangkaty():real;
begin 
    Write('Masukan angka x : ');
    readln(x);
    write('Masukan Pangkat y : ');
    readln(y);
    index:=1;
    pangkat:=1;
    while index<=y do
    begin  
    pangkat:=pangkat*x;
    index+=1;
    end;
    Writeln(pangkat);
end;

function barisgeo():real;
begin
  write('masukan angka awal (a) = ');readln(a);
  write('masukan jarak (r) = ');readln(r);
  write('masukan batas deret (n) = ');readln(n1);

  index:=1;
  
while (index<=n1)do
  begin
    a:=a*r;
    index+=1;
  end;
  writeln(a);
end;




begin
clrscr();
//barisgeo();
//xpangkaty();
faktorialn();
//kpk();

  
  ReadLn();
end.
