program Pertemuan3;
uses crt;
var
  angka,baris,kolom :integer;
  i, hasil,hasil1,hasil2:Integer;
  bagi : real;
  j: Integer;
  
function satu():integer;
begin
    write('Input Sebuah Angka : ');
    ReadLn(angka);
    for i:=1 to 10 do
    begin
      hasil := angka mod i;
      bagi := angka/i;
      WriteLn(angka,':',bagi:0:0,'sisa', hasil);
    end;
end;


function dua():Integer ;
begin  
    write('Input angka faktorial : ');
    ReadLn(angka);
    hasil:=1;
    for i:=1 to angka do
      begin
      hasil := hasil*i;
      end;
      write(angka,' Faktorial = ',hasil);
end;

function tiga():integer;
begin
    write('input batas = ');
    readln(angka);
    for i:=1 to angka do
    begin
      if (i mod 3 = 0) or (i mod 5 = 0) and (i < angka) then
      begin
        write(i,' ');
        if (i mod 2 = 0) then
        hasil1 := hasil1+1
        else if (i mod 2 <>0 ) then
        hasil2+=1;
       end;
    end;
    writeln('jumlah angka genap = ',hasil1);
    writeln('Jumlah Angka Ganjil = ',hasil2);
end;



function empat():integer;
begin
write('baris = ');
readln(baris);
write('kolom = ');
readln(kolom);

i:=1;
   while i <= baris do
   begin
   j:=1;
      while j <= kolom do
      begin
      write(j);
      j+=1;
      end;
      writeln;
     i+=1;
   end;
end;

begin
  //satu();
  //dua();
  //tiga();
  empat();
  readln;
end.

