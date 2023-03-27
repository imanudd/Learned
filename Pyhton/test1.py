print("Masukan Angka : ")
speed = int(input())

def DefineSpeed(int):
    if speed >= 1000:
        print("extremly fast")
        return
    if speed >= 150:
        print("fast")
        return
    if speed >= 50:
        print("normal")
        return
    if speed >= 10:
        print("average") 
    else: print("too slow")


DefineSpeed(speed)




