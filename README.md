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



results on 24 physical cores, 256G RAM  


	1 million files 

		go
			Filter took     60.300411ms, 16583.634894 files/ms
			Map took        274.672488ms, 3640.699537 files/ms

			Filter took     96.926585ms, 10317.086896 files/ms
			Map took        286.672647ms, 3488.299321 files/ms

			Filter took     91.170182ms, 10968.498451 files/ms
			Map took        273.915927ms, 3650.755219 files/ms

			Filter took     50.224714ms, 19910.516564 files/ms
			Map took        191.739631ms, 5215.405885 files/ms

			Filter took     108.574928ms, 9210.229456 files/ms
			Map took        189.974183ms, 5263.873144 files/ms

		scala
			Filter took:  54.30776 ms, 18413.57422 files/ms
			Map took:     1049.86328 ms, 952.50500 files/ms

			Filter took:  71.46753 ms, 13992.36816 files/ms
			Map took:     93.31990 ms, 10715.82813 files/ms

			Filter took:  5.15910 ms, 13305.10840 files/ms
			Map took:     1895.46802 ms, 527.57416 files/ms

			Filter took:  76.12512 ms, 13136.26758 files/ms
			Map took:     93.26137 ms, 10722.55371 files/ms

			Elapsed time: 78.35020 ms, 12763.20898 files/ms
			Map took:     98.85263 ms, 10116.06836 files/ms

			
	10 million files
		go	
			Filter took     504.807036ms, 19809.549564 files/ms
			Map took        2.75167996s, 3634.143558 files/ms
			
			Filter took     514.307152ms, 19443.633947 files/ms
			Map took        2.715162621s, 3683.020650 files/ms

			Filter took     554.459ms, 18035.598665 files/ms
			Map took        2.600349589s, 3845.636772 files/ms


			Filter took     1.531389489s, 6530.017394 files/ms  (yikes)
			Map took        2.787230807s, 3587.790424 files/ms

			Filter took     544.849088ms, 18353.706045 files/ms
			Map took        2.843270805s, 3517.076172 files/ms

		scala
			out of memory error? 
