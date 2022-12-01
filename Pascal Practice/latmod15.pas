program latmod15;
uses crt;
type
  array_angka=array[1..5]of Integer ;
  
var 
  data_angka:array_angka;
  i:Integer;
  max,min:Integer;
  
function input_array():integer;
begin
for i:=1 to 5 do
  begin
      write('input angka ke ', i ,'= ');
      readln(data_angka[i]);
  end;
end;

function show_array(var x : array_angka):integer;
begin
for i:=1 to 5 do
  begin
      write(data_angka[i],' ');
  end;
end;

function maximum(var m:array_angka):integer;
begin
for i:=1 to 5 do
begin
   if (i = 1) then
      max := m[i]
    else if (m[i] > max) then
      max := m[i];
end;
writeln('maximum= ',max);
end;


function minimum(var m:array_angka):integer;
begin
for i:=1 to 5 do
begin
   if (i = 1) then
      min := m[i]
    else if (m[i] < min) then
      min := m[i];
end;
writeln('minimum = ',min);
end;


begin
  clrscr();
  input_array();
  show_array(data_angka);
  writeln;
  maximum(data_angka);
  minimum(data_angka);
  
  
  
  
  readln;
end.
