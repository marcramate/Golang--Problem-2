<!-- 1.func maxProChickens(n, k int, poschickens []int) int
    n (จำนวนไก่), k (ความยาวหลังคา), poschickens (ตำแหน่งของไก่)
2. sort.Ints(chickens) // เรียงลำดับตำแหน่งไก่ให้เป็นลำดับเพิ่มขึ้น
left, right := 0, 0 // ตั้งค่าเริ่มต้นของ left และ right pointer
maxChickens := 0 // ตั้งค่าเริ่มต้นของจำนวนไก่สูงสุดที่สามารถปกป้องได้    

for right < n { // วนลูปจนกระทั่ง right ถึงท้ายสุดของ chickens
		end := chickens[right] + k // คำนวณจุดสิ้นสุดของช่วงที่หลังคาปกป้องได้
		chickensInRange := 0 // นับจำนวนไก่ในช่วงนี้

	for left < n && chickens[left] < end { // เลื่อน left pointer ไปข้างหน้าเพื่อนับจำนวนไก่ในช่วงปกป้อง
			chickensInRange++
			left++
		}

        maxChickens = max(maxChickens, chickensInRange) // อัปเดตค่า maxChickens ให้เป็นค่ามากสุด
		right++ // เลื่อน right pointer ไปข้างหน้า
	}

    	return maxChickens // คืนค่าจำนวนไก่สูงสุดที่สามารถปกป้องได้


        

 -->