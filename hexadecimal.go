package main
import ("fmt")

func main(){
	adams:= 43
	fmt.Printf("43 as binary, %b \n", adams)
	fmt.Printf("43 as hexadecimal, %X \n", adams)

    a , b, c, d, e, f := 0 , 1, 2, 3, 4, 5
	fmt.Printf("%v \t %b \t %#X\n",a,a,a)
	fmt.Printf("%v \t %b \t %#X\n",b,b,b)
	fmt.Printf("%v \t %b \t %#X\n",c,c,c)
	fmt.Printf("%v \t %b \t %#X\n",d,d,d)
	fmt.Printf("%v \t %b \t %#X\n",e,e,e)
	fmt.Printf("%v \t %b \t %#X\n",f,f,f)	
		
}
