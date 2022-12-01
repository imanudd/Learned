program latihansoal6;
uses crt;
var
  gapok,tunj: Integer;
  gol_pend,gol_kar,command:String;

function header():string;
begin
  WriteLn('--------------------------------');
  WriteLn('PROGRAM PERHITUNGAN GAJI PEGAWAI');
  WriteLn('--------------------------------');
end;

function footer():string;
begin
  WriteLn('--------------------------------------------');
  WriteLn('TERIMA KASIH TELAH MENGGUNAKAN PROGRAM KAMI!');
  WriteLn('--------------------------------------------');
end;

function inputgol_kar():string;
begin 
while (gol_kar <> 'A') And (gol_kar <> 'B') do
    begin
      write('masukan golongan Karyawan (a/b) = ');
      readln(gol_kar);
      gol_kar := upcase(gol_kar);
      if (gol_kar <> 'A') and (gol_kar <> 'B') then
      begin
      WriteLn('**************************************');
      WriteLn('Salah Memasukan Golongan Karyawan, mohon ulangi!');
      WriteLn('**************************************');
      end;
    end;
end;

function inputpend():string;
begin
while (gol_pend <> 'SMA') And (gol_pend <> 'S1') do
    begin
      write('masukan golongan Pendidikan (SMA/S1) = ');
      readln(gol_pend);
      gol_pend:=upcase(gol_pend);
      if (gol_pend <> 'SMA') and (gol_pend <> 'S1') then
      
      begin
      WriteLn('**************************************');
      WriteLn('Salah Memasukan Golongan Pendidikan mohon ulangi!');
      WriteLn('**************************************');
      end;
      
    end;
end;

function gaji(var x,y:Integer):integer;
  begin
    gaji := x+y;
  end;
  
function penentuan_gaji(var gol,pend:String):string;
  begin
    if (pend = 'SMA') then 
    begin
      if (gol= 'A') then
        begin
         gapok:=3000000;
         tunj:=2000000;
         WriteLn('Gaji total adalah Rp.',gaji(gapok,tunj));
        end;
        if (gol= 'B') then
        begin
         gapok:=4000000;
         tunj:=3000000;
         WriteLn('Gaji total adalah Rp.',gaji(gapok,tunj));
        end;
    end
    else 
    if (pend = 'S1') then
      begin
       if (gol = 'A') then
        begin
         gapok:=4000000;
         tunj:=3000000;
         WriteLn('Gaji total adalah Rp.',gaji(gapok,tunj));
        end
       else
        begin
         gapok:=5000000;
         tunj:=4000000;
         WriteLn('Gaji total adalah Rp.',gaji(gapok,tunj));
        end;
      end;
end;

function loop():string;
begin 
repeat
begin
      Write('Apakah anda ingin melanjutkan? (y/n)');
      readln(command);
      command:=UpCase(command);
      writeln('-----------------------------',command);
    end;
    until (command = 'Y') or (command = 'N');
end;


begin
repeat
begin
  clrscr;
  header();
  inputgol_kar();
  inputpend();
  writeln;
  writeln('Karyawan dengan golongan ',gol_kar,
  ' dengan pendidikan ', gol_pend);
  penentuan_gaji(gol_kar,gol_pend);
  writeln;
  footer();
  loop();
end;
until command='N';
  readln;
  
  
  
end.


