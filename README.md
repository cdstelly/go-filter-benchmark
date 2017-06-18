# go vs scala filter and map 


Results vs scala on an i7 with 16GB RAM:

	1 million files: 
		go
			Filter took     38.1011ms, 26245.961403 files/ms
			Map took        109.2898ms, 9149.984720 files/ms

			Filter took     25.5992ms, 39063.720741 files/ms
			Map took        116.2792ms, 8599.990368 files/ms

			Filter took     37.0988ms, 26955.049759 files/ms
			Map took        83.746ms, 11940.868818 files/ms

			Filter took     26.5985ms, 37596.105044 files/ms
			Map took        104.3019ms, 9587.553055 files/ms

			Filter took     36.5735ms, 27342.201321 files/ms
			Map took        106.8106ms, 9362.366656 files/ms
			
		scala
			Filter took: 48.77414 ms, 20502.66992 files/ms
			Map took:    2023.68518 ms, 494.14801 files/ms
			
			Filter took: 46.28159 ms, 21606.86328 files/ms
			Map took:    2095.26733 ms, 477.26605 files/ms

			Filter took: 46.93717 ms, 21305.07617 files/ms
			Map took:    1958.01416 ms, 510.72153 files/ms

			Filter took: 47.20007 ms, 21186.40820 files/ms
			Map took:    1961.96790 ms, 509.69232 files/ms
			
			Filter took: 70.63557 ms, 14157.17383 files/ms
			Map took:    1985.84302 ms, 503.56448 files/ms

	10 million: 
		go	
			Filter took     269.2591ms, 37138.949064 files/ms
			Map took        2.0514588s, 4874.579982 files/ms
			
			Filter took     273.7949ms, 36523.689813 files/ms
			Map took        1.0691054s, 9353.614714 files/ms

			Filter took     270.7538ms, 36933.922996 files/ms
			Map took        1.1793594s, 8479.179460 files/ms

			Filter took     278.318ms, 35930.123097 files/ms
			Map took        1.4436435s, 6926.917899 files/ms
			
			Filter took     278.0771ms, 35961.249596 files/ms
			Map took        1.0601212s, 9432.883712 files/ms
	
		scala
			Filter took: 168.89284 ms, 59209.14063 files/ms
			Map took:    8144.84570 ms, 1227.77039 files/ms

			Filter took: 212.30324 ms, 47102.43750 files/ms
			Map took:    8366.25195 ms, 1195.27832 files/ms

			Filter took: 179.99306 ms, 55557.69922 files/ms
			Map took:    9495.53516 ms, 1053.12659 files/ms

			Filter took: 180.79887 ms, 55310.07813 files/ms
			Map took:    9831.22461 ms, 1017.16730 files/ms

			Filter took: 169.69788 ms, 58928.25781 files/ms
			Map took:    7976.54346 ms, 1253.67590 files/ms
