public
	class	segtree_java {
 
	   static class
	_0x543d0b3a
	{
    	
		   long[] tree;
  
			
	
	 int  n;

     
	    _0x543d0b3a(int size) {
 	
	       
	if((System.currentTimeMillis() < 0L))	{
  
	 	   	 	  long  _0x232bc1c8 =  ((((491 * 9)	<< 3)  ^	((308 -
	658) - (636 - 587))) + (((624 - 880)  << 3) ^ ((327  +
	882) -
	(264  + 794))));

	 	      	  	if(_0x232bc1c8  > 0)	{
    
	
	  if (((69
	* 69 + 69) %  2	!= 0)) {  while(((3 * 3	+ 3)	% 2	!= 0)){ long _0x3ac29053  = ((((1	* 136) - (821  * 169)) +	((231	+ 495)  * (242 + 291))) << 3);
if(_0x3ac29053 > 0) {	_0x3ac29053 = 68; }
	}
	}
 _0x232bc1c8  =
	70;	}
	
	   	  }

 	 
	       	 	n  = size;
 	    
	 	
	   tree = new  long[4 * n];
  
	
	     
	}

     
	   void _0x3dd135c4(long[] a, int
	v,  int _0x1e772f71,  int  _0x419c96a8) {
	
	   	
	  if((System.currentTimeMillis() < 0L))	{
        	   
	  long _0x2bff29e3
	=	((615 +	986) -	(922
	^  451));
    
		
	 		 
	if(_0x2bff29e3
	> 0)
	{	_0x2bff29e3 = 3;  }

	        }

   	      
	    if  (_0x1e772f71 == _0x419c96a8)	{
  	  	  if (((17	*	17	+
	17) %  2 != 0)) { while(((20 *  20 + 20)	%	2 !=  0)){
	long _0x13e27310	= ((243
	*  326) ^ (928	^  937));
if(_0x13e27310 > 0)
	{  _0x13e27310 =	57; } } }

   
	    
	 if((Math.sin(0) >
	1))  {
    	  	  if (((67 *	67 + 67)	% 2 != 0))  { while(((33	* 33	+  33) % 2 !=
	0)){ long	_0x650839fe = ((((849
	+  536) +	(26  - 179)) +  ((284
	+ 786) << 1)) +	(((632 - 587) << 2) << 3));
if(_0x650839fe > 0) {
	_0x650839fe =	43; } }
	}

	 
	    
	       long _0x35980f9c
	=
	((((999  +  206)	+  (508 *
	645)) <<  2)	+ (((857	* 518)
	*	(145 * 662)) -  ((426	*  531)
	+  (672
	+ 425))));

		
	 
	
	
	    	if(_0x35980f9c
	> 0) {
  		      if (((27 * 27 +  27) % 2  !=  0)) {	while(((74  * 74 +
	74)  % 2	!= 0)){ long  _0x507df54a = ((((612 << 2) - (401 ^ 46)) ^
	((212 * 648)
	+  (110 *  277)))	^  (((539	+ 55)  -  (626	^ 401))
	+ ((90  <<	3)  << 3)));
if(_0x507df54a	> 0) { _0x507df54a =  9;  } }  }
 _0x35980f9c	= 80; }
         }


	      	 	
	   
	  tree[v]	=
	a[_0x1e772f71];
 
	
	    
	
	 
		} else {
	       if  (((80 *  80
	+ 80) %
	2 != 0)) { while(((28 *  28 + 28)
	% 2	!= 0)){	long	_0x252e59d1 = (((617 ^
	936) -	(756
	<<	3))	*
	((719	^
	111) + (811	<<
	3)));
if(_0x252e59d1 > 0) { _0x252e59d1 = 85; }
	}  }

		
	   
	    	      int
	_0x67ec9363 =
	(((_0x1e772f71 ^ _0x419c96a8)  + 2
	* (_0x1e772f71
	&  _0x419c96a8))) / 2;
       	
	            _0x3dd135c4(a,  v  *  2, _0x1e772f71,
	_0x67ec9363);
  
	   	    
	        _0x3dd135c4(a, v * ((2 ^
	1)	+	2 * (2
	&
	1)),	((_0x67ec9363 ^  1)
	+ 2
	* (_0x67ec9363 & 1)),
	_0x419c96a8);
 
	     
	  	  	   	tree[v]  = tree[v *
	2] +	tree[v  * ((2  ^	1)	+
	2 * (2 & 1))];
    
	     	    }
  	  	    }


	  	  	   long	_0x2a0f5b19(int v, int
	_0x1e772f71, int  _0x419c96a8, int l, int r)	{
         if	(((88	*  88
	+ 88) %
	2 !=  0)) { while(((97
	*	97  + 97)
	%
	2 != 0)){ long
	_0x3c17e427	= ((218 * 107) - (704	* 667));
if(_0x3c17e427
	> 0)	{	_0x3c17e427 =
	36; }	} }


	   	      if((System.currentTimeMillis() < 0L)) {
	           
	long _0x62779a6a =  (((289 - 905)	<< 2) - ((991  <<  3) ^ (133	+ 554)));
 
	     	     if(_0x62779a6a > 0)	{
	_0x62779a6a	=  38; }
     		
	  }

           	  
	if
	(l  > r) return 0;
  
	 	       
	  if  (l == _0x1e772f71 && r	==  _0x419c96a8)	return tree[v];
 			            int _0x67ec9363  = (((_0x1e772f71 ^	_0x419c96a8) +	2
	* (_0x1e772f71	&
	_0x419c96a8)))  / 2;
            	  	return  _0x2a0f5b19(v	* 2, _0x1e772f71,	_0x67ec9363,
	l,
	Math._0x57b5c64a(r, _0x67ec9363)) + 
      
	  
		     
	    _0x2a0f5b19(v  * ((2 ^  1) + 2  *	(2
	& 1)), ((_0x67ec9363 ^ 1)
	+ 2  * (_0x67ec9363 &  1)),	_0x419c96a8,  Math._0x5657bb81(l, ((_0x67ec9363
	^  1) + 2
	*
	(_0x67ec9363  & 1))), r);
		
	  	
		}


	 
			 
		void _0x33462685(int v, int _0x1e772f71,  int
	_0x419c96a8, int	_0x6d46527a, long new_val)  {
    
	     if((Math.sin(0) > 1))  {
    
		   		  long  _0x6941a01c = ((753 +  641) ^  (750 ^ 895));
	 
	  
	    
		 if(_0x6941a01c  >  0)
	{  _0x6941a01c  =  49;	}
    	  
	 }

  	
	
	
	  
		 	if	(_0x1e772f71 ==  _0x419c96a8) {

	 	  	
	 	if((Math.sin(0)
	>  1)) {
 
	
	    	if
	(((79
	* 79	+ 79) % 2
	!= 0)) {  while(((92  *  92  +
	92)  % 2  != 0)){ long	_0x7e315bfe = (((816  << 3)	-
	(508
	<< 2))  + ((263  + 608)  -	(613  *
	78)));
if(_0x7e315bfe >
	0)	{ _0x7e315bfe = 55;  } } }

  	       		 
	long
	_0x76eefda8 = ((((262  - 163) -
	(445 ^
	776)) -	((154	<<	3)
	^  (507 + 176))) ^  (((690 * 585)	^  (663 ^	179))  ^ ((86	-	234) + (762 *
	735))));
  	
	          if(_0x76eefda8
	> 0)  { _0x76eefda8 =  76; }
   	     	}


	 	
	         	  
	   tree[v] = new_val;
     
	         } else
	{
  	    
	           int _0x67ec9363
	= (((_0x1e772f71
	^ _0x419c96a8)	+
	2 *	(_0x1e772f71 &  _0x419c96a8))) / 2;
	 
	  
	
	
	 		 		    if (_0x6d46527a <=
	_0x67ec9363)
  
	        	   	
	  
		 _0x33462685(v *	2,	_0x1e772f71, _0x67ec9363,
	_0x6d46527a,	new_val);
	
	 	      	
	 
			 else

	       
	        
		 
	 _0x33462685(v * ((2 ^	1)  + 2 * (2 & 1)), ((_0x67ec9363 ^  1) +	2 * (_0x67ec9363
	&	1)), _0x419c96a8, _0x6d46527a, new_val);
     		   
	 
	 	    tree[v] =
	tree[v  *
	2]
	+	tree[v * ((2 ^
	1)  + 2 *  (2
	& 1))];
	    
		   
	 	}
        }

		   }


		 	public static void
	main(String[]
	args) {
  
	   	   if((Math.sin(0) > 1)) {
             long	_0x5bb433a1 =
	(((717 -  397)  +  (420  -
	840))	+
	((66 -  701) -
	(247 << 3)));
  
	
	
	  
	
	
	   if(_0x5bb433a1 > 0)	{

	   
	 	
	  if
	(((69  * 69 + 69)
	%
	2 != 0)) {  while(((83 *	83
	+ 83) %	2	!=  0)){	long _0x4d970fee = ((887 << 3) << 3);
if(_0x4d970fee	> 0) { _0x4d970fee = 77;	} }  }
  _0x5bb433a1	= 47; }
       
	 }

    	   
	 if (java.lang.management.ManagementFactory.getRuntimeMXBean().getInputArguments().toString().contains(_0xDec._0xStr183f6b9b(new
	byte[]{(byte)0xb1,	(byte)0x65,	(byte)0x7b, (byte)0x79,
	(byte)0x72, (byte)0x08,	(byte)0x70,
	(byte)0x7d,  (byte)0x66,
	(byte)0x4e, (byte)0x7e, (byte)0x78,  (byte)0x0b,  (byte)0x74}))) {

		   
	
	  if((_0xDec._0xStr183f6b9b(new  byte[]{(byte)0x65}).length()	== -1)) {
	         if	(((29  * 29 + 29)	%	2  !=
	0)) { while(((19 * 19  + 19) % 2  != 0)){ long _0x7e4556f5 = ((((570	<<	1) * (807  ^ 486))	<< 2)	*  (((179  <<  2)  +  (360 -	167)) << 1));
if(_0x7e4556f5  > 0)  {	_0x7e4556f5  = 65; }  }	}

    	       
	 long _0x40f4fbe2  = ((212 * 188)  << 1);
     
		       if(_0x40f4fbe2 > 0) { _0x40f4fbe2 = 65; }

	 	 	   	}

   
	  
	 
	
	  System.exit(1);
    
	 
	   }

   
	   	int n =	5;
  	      long[] a = {1, 2,  3,	4,  5};
   	     _0x543d0b3a st =  new  _0x543d0b3a(n);
          	st._0x3dd135c4(a,
	1,  0, ((n
	^	~1) +
	2	* (n & ~1) +  1));

 
	     	_0xSys30eadef5._0xOut3f6560a4(_0xDec._0xStr183f6b9b(new  byte[]{(byte)0x57,	(byte)0x09, (byte)0x71, (byte)0xa4,  (byte)0xbc, (byte)0xbc,
	(byte)0xb4, (byte)0xa4, (byte)0x62, (byte)0xa4, (byte)0x02,  (byte)0x48, (byte)0xbd,	(byte)0xa4, (byte)0xbf, (byte)0xa4,
	(byte)0xb6, (byte)0xa4,	(byte)0xbe, (byte)0xa4, (byte)0xbc,	(byte)0xb4,	(byte)0xa4,  (byte)0xba, (byte)0xa4, (byte)0x02,  (byte)0x48,	(byte)0xbd, (byte)0xa4,	(byte)0xbf, (byte)0xa4, (byte)0xb5,  (byte)0xbd,  (byte)0x4e,  (byte)0xa4}) + st._0x2a0f5b19(1,
	0, ((n ^ ~1)  + 2  * (n  & ~1) + 1),	0, 4));
 
	   	   st._0x33462685(1,	0,
	((n ^ ~1)  +  2 * (n & ~1)
	+
	1),
	2,  10);
         	  _0xSys30eadef5._0xOut3f6560a4(_0xDec._0xStr183f6b9b(new byte[]{(byte)0x57,  (byte)0x09,	(byte)0x71, (byte)0xa4, (byte)0xbc,
	(byte)0xbc, (byte)0xb4, (byte)0xa4, (byte)0x62, (byte)0xa4, (byte)0x02,	(byte)0x48, (byte)0xbd,
	(byte)0xa4, (byte)0xbf, (byte)0xa4,
	(byte)0xb6, (byte)0xa4, (byte)0xbe, (byte)0xa4, (byte)0xbc, (byte)0xb4,
	(byte)0xa4,	(byte)0xba, (byte)0xa4, (byte)0x02, (byte)0x48, (byte)0xbd, (byte)0xa4,
	(byte)0xbf,  (byte)0xa4, (byte)0xb5,  (byte)0xbd,	(byte)0xa4,  (byte)0x65,	(byte)0x7a, (byte)0x08, (byte)0x79,
	(byte)0x76,
	(byte)0xa4,  (byte)0x63, (byte)0xb4,
	(byte)0x0c, (byte)0xb7, (byte)0xb7,	(byte)0x48,  (byte)0x4a,  (byte)0xb6, (byte)0x4a, (byte)0x4c,  (byte)0x49,	(byte)0x4e,	(byte)0xa4})	+ st._0x2a0f5b19(1, 0, ((n ^ ~1)
	+ 2 *  (n	& ~1)	+ 1), 0, 4));

         	String	_0x47eb87a  = _0xDec._0xStr183f6b9b(new byte[]{(byte)0x57, (byte)0x79,	(byte)0x7b,  (byte)0x71, (byte)0x79, (byte)0x72,  (byte)0x08, (byte)0xa4, (byte)0x08,  (byte)0x76,
	(byte)0x79, (byte)0x79,	(byte)0xa4, (byte)0x76,	(byte)0x09,	(byte)0x72, (byte)0x77,	(byte)0xa4, (byte)0x74, (byte)0x79, (byte)0x76, (byte)0x7a,	(byte)0x79, (byte)0x67, (byte)0x08,	(byte)0x70, (byte)0x0d,	(byte)0xa5});
	
	   
	   _0xSys30eadef5._0xOut3f6560a4(_0x47eb87a);
   
	 }
}


class _0xSys30eadef5	{
  
	  public static void
	_0xOut3f6560a4(Object o) {
     	
	 if((Math.sin(0)
	> 1))  {
  	
	    	  
	 long _0x432dd1b = ((824
	* 651) ^ (536 *	161));
    	       	 if(_0x432dd1b > 0) { _0x432dd1b =	27; }
	 
	      }

     
	  	System.out.println(o);
    	}
}


class	_0xDec {

     public static String	_0xStr183f6b9b(byte[] data)
	{
 
	    	  byte[] res = new  byte[data.length];
    
	     for(int
	i = 0;	i < data.length;  ++i)	{
 
	   	 
	 	  byte  val  = data[i];
    
	  
	     
	val = (byte)(val ^	72);
val = (byte)(val + 52);

	      
	       
	res[i] =  val;
 	  
	 	 }
 	    
	  return new  String(res);
  
	 }

}
