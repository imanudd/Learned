program caseoflatsol7;
uses crt;
var
  input:integer ;
  command:String ;
  cekinput:integer;
  
  
function header():string;
begin
  clrscr;
  WriteLn('--------------------------------');
  WriteLn('   PROGRAM PENCARI NAMA BULAN   ');
  WriteLn('--------------------------------');
end;

function footer():string;
begin
  WriteLn('--------------------------------------------');
  WriteLn('TERIMA KASIH TELAH MENGGUNAKAN PROGRAM KAMI!');
  WriteLn('--------------------------------------------');
end;

function loop():string;
begin
repeat 
    Begin
      Write('Apakah anda ingin melanjutkan? (y/n)');
      readln(command);
      command:=upcase(command);
      writeln('-----------------------------');
    end;  
  until (command = 'Y') or (command = 'N');
end;

function input_bulan():integer;
begin
repeat
   begin
   Write('Masukan Angka Untuk Menentukan Bulan (1-12) = ');
   readln(input);
   if input > 13 then
    begin
    writeln('xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx');
    writeln('  Salah Memasukan Input, Ulangi!  ');
    writeln('xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx');
    end;
   end;
until input <= 12;
end;

function caseof(var inputan:integer):string;
begin
  case(inputan)of
  1 : writeln('January');
  2 : writeln('February');
  3 : writeln('Maret');
  4 : writeln('April');
  5 : writeln('Mei');
  6 : writeln('Juni');
  7 : writeln('Juli');
  8 : writeln('Agustus');
  9 : writeln('September');
  10: writeln('Oktober');
  11: writeln('November');
  12: writeln('Desember');
  end;
end;

begin

  repeat
  begin
  header();
  input_bulan();
  writeln;
  writeln;
  writeln('+++++++++++++++++++++++++++++++++++++');
  write('BULAN YANG ANDA MAKSUD ADALAH ');
  caseof(input);
  writeln('++++++++++++++++++++++++++++++++++++++');
  writeln;
  writeln;
  footer();
  loop();
  end;
  
  until command = 'N'; 
  readln;
  
  
end.
