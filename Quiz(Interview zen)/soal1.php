$n = 10;
for ($i=1;$i<=$n;$i++){
    for($j=1;$j<=$i;$j++){
        if($j==1 or $j==$a){
            echo "*";
            $a=$j+3;
        }else if($j==2 or $j==$b){
            echo "#";
            $b=$j+3;
        }else if($j==3 or $j==$c){
            echo"%";
            $c=$j+3;
        }
    }
    echo"\n";
}