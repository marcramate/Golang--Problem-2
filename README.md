# Golang Problem 2: Superman's Chicken Rescue
### Description:
* English Version :
  * In a one-dimensional world, Superman needs to protect chickens from a heavy rainstorm using a roof of limited.
* Thai Version :
  *  ในอาเรย์มิติเดียว(โลก1มิติ)  ซูเปอร์แมนจำเป็นต้องปกป้องได้จากพายุฝนโดยใช้หลังค่าที่มีความยาวจำกัด
## Use Language
 * Golang 
## EX.
* Englist Version : 
   * Create Function  maxProChickens and used 'sort' module sort the position chickens and find end point that a roof covers. Replace 'n' represents total chickens, 'k' with the lenght and 'poschicken' with position chickens.
Process loop each position until you reach the position last chicken and agin loop right until position last chicken and loop left until last position chicken but beyond the position chicken from the right .
 * Thai Version :
   * สร้างฟังก์ชั่น maxProChickens แล้วใช้ module 'sort' เข้ามาช่วยในการเรียงลำดับ และทำการนับตำแหน่งของไก่และหาจุดสิ้นสุดที่หลังคาคลุมถึง โดยจะมีการแทน n ด้วยจำนวนไก่ทั้งหมด k ด้วยความยาว และ poschicken ด้วยตำแหน่งของไก่แต่ละตัว และทำการลูปไปที่ละตำแหน่งจนกว่าจะถึงตำแหน่งของไก่ตัวสุดท้าย
ที่ซูเปอร์แมนสามารถปกป้องได้ โดยการนับไปทางขวาจนกว่าจะถึงตำแหน่งสุดท้ายของไก่ และทำอีกครั้งโดยการเลื่อนไปท้ายซ้ายจนถึงถึงตำแหน่งของไก่ตัวสุดท้ายแต่ไม่เกินตำแหน่งของไก่จากทางขวา
```
//Example 1
n1, k1 :=  5, 5
poschickens1 := []int{2, 5, 10, 12, 15}
```

```
//Example 2
n2, k2 :=  6, 10
poschickens2 := []int{1, 11, 30, 34, 35, 37}
```
## Result
> Example 1
```
 Result : Maximum chickens : 2
 ```
>  Example 2
```
 Result : Maximum chickens : 4
 ```
