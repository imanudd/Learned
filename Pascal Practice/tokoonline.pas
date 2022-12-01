program tokoonline;
uses crt;

var
  product:integer;
  idproduct:String;
  qty,ongkir:Integer;
  bonus,dis,ppn:real;
  dsc,total,finaltotal:real;
  command,delivery:string;
  
function listbarang():String ;
begin
  clrscr();
  writeln('-------------------------------------------');
  writeln(' ++++++++++++++++LIST BARANG+++++++++++++++');
  writeln('-------------------------------------------');
  writeln('+ 1. GL01 | Gula Pasir         | Rp.12.000 ');
  writeln('+ 2. MG01 | Minyak Goreng      | Rp.13.500 ');
  writeln('+ 3. BR01 | Beras Pandan Wangi | Rp.10.000 ');
  writeln('+ 4. GS01 | Gas Elpiji 12kg    | Rp.72.000 ');
  writeln('-------------------------------------------');
end;



function inputproduct():integer;
begin
idproduct:='';
while (idproduct <> 'GL01') and (idproduct <> 'MG01')
      and (idproduct <>'BR01') and (idproduct <> 'GS01') do      
      begin
      write('Masukan Kode Productnya = ');
      readln(idproduct);
      idproduct := upcase(idproduct);
      
      if (idproduct = 'GL01')then
      product:=12000
      else if (idproduct = 'MG01')then
      product:=13500
      else if (idproduct = 'BR01')then
      product:=10000
      else if (idproduct = 'GS01')then
      product:=72000
      
      Else
      begin
      WriteLn('**************************************');
      WriteLn('Salah Memasukan Kode Product, mohon ulangi!');
      WriteLn('**************************************');
      end;
    end;
end;



function input_qty():integer;
begin
    write('masukan jumlah yg ingin dibeli = ');
    readln(qty);
end;



function kalkulasi(var prdct:string ; var quantity:integer):integer;
begin
   if (prdct='BR01')and (quantity >= 5) then
      begin
      writeln('Selamat anda Mendapatkan Bonus Minyak Goreng');
      dsc:= 5/100;
      end
   else if (quantity>=5) then
      dsc:=2/100
   else
     writeln('false');
end;



function input_delivery():string;
begin
while (delivery <> 'YES') And (delivery <> 'NO') do
    begin
      write('Apakah Pesanan Ingin Di antar?(YES/NO) = ');
      readln(delivery);
      delivery:=upcase(delivery);
      if (delivery <> 'YES') and (delivery <> 'NO') then
      
      begin
      WriteLn('**************************************');
      WriteLn('Salah Memasukan inputan,  mohon ulangi!');
      WriteLn('**************************************');
      end;
      
    end;
end;



function total_bayar(var dlv:string; var dsc:real;var product:integer): real ;
begin
   total:=product*qty;
   ppn := total*10/100;
   ongkir := 5000;
   dis := total*dsc;
   if (dlv='YES') then
      finaltotal := (total + ongkir + ppn)-dis
    else
    begin
    ongkir:=0;
    finaltotal :=(total + ongkir + ppn)-dis;
    end;
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
begin

  repeat 
  begin
  listbarang();
  inputproduct();
  input_qty();
  input_delivery();
  kalkulasi(idproduct,qty);
  total_bayar(delivery,dsc,product);
  writeln('------------------------------');
  writeln('SubTotal Belanjaan = ',total);
  writeln('Diskon Belanjaan = ',dis);
  writeln('Delivery Fee = ',ongkir);
  writeln('PPN Sebesar = ',ppn);
  writeln('------------------------------');
  writeln('Total nya adalah = ',finaltotal);
  loop();
  
  end;
  until command = 'N';
  readln;
end.
