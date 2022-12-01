program MuhammadImanuddinJamil_BMI;
uses crt;
var data:Integer;
    jumlah,j_data:integer;

function identitas():String;
begin
  clrscr();
  WriteLn('Nama  : Muhammad Imanuddin Jamil');
  WriteLn('Npm   : 202243501459');
  WriteLn('Kelas : XL');
end;

function footer():string;
begin
  WriteLn('--------------------------------------------');
  WriteLn('TERIMA KASIH TELAH MENGGUNAKAN PROGRAM KAMI!');
  WriteLn('--------------------------------------------');
end;

function methodrepeat():integer;
begin
repeat
  writeln('masukan data = ');
  readln(data);
  jumlah := jumlah+data;
  j_data+=1;
  until jumlah >= 50;
  writeln('---------------------------------');
  Write('Banyak nya data = ',j_data);
  writeln(' Jumlahnya ',jumlah);
end;

begin
  identitas();
  writeln;
  methodrepeat();
  footer();
  readln;
end.
