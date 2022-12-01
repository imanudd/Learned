program NestedIF;
uses crt;
var
  nilai1, nilai2, nilai3: Integer;
  command : String;
  pembanding : integer ;
begin 
 
  writeln('OPERATOR PERBANDINGAN 3 ANGKA');
  WriteLn('=============================');
  
  command:='';
  while (command <> 'y') and (command <> 'n') do
  Begin
   Write('Apakah anda ingin melanjutkan? (y/n)');
  readln(command);
  end;
  
  while (command <> 'n')
     do 
    begin
      WriteLn('=======================');
      Write('Masukan Nilai A = ');
      readln(nilai1);
      Write('Masukan Nilai B = ');
      readln(nilai2);
      Write('Masukan Nilai C = ');
      readln(nilai3);
      if nilai1=nilai2 and nilai3 then
      writeln('Tidak Dapat Mengoprasikan 3 Angka Yang Sama')
      else
        if (nilai1 > nilai2) AND (nilai1 > nilai3) then
            WriteLn('Nilai Yang Terbesar Adalah : ',nilai1)
        
            else
            if (nilai2 > nilai1) AND (nilai2 > nilai3) then
               writeln('Nilai Yang Ter terbesar adalah : ',nilai2)
   
            else
            WriteLn('Nilai Yang Terbesar Adalah : ',nilai3);
            Writeln;
            
    command:='';
    while (command <> 'y') and (command <> 'n') do
    Begin
    writeln('----------------------------------');
    Write('Apakah anda ingin melanjutkan? (y/n)');
    readln(command);
    writeln('----------------------------------');
    end;
end;

end.
