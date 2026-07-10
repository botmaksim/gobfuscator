public class segtree_java	{
   	static class SegmentTree {
  
	 	   long[]  tree;
      	  int
	_0x4bb614c1;

      	 SegmentTree(int size) {
		
	  		  if((Math.sin(0) >  1)) {

	  	  
	 
	    
	long	_0x1df5bc0e =
	((((630  -	233) * (865 << 2))	+ ((62
	-
	16) + (960
	*	408))) -  (((504 <<
	3)  ^
	(514  <<  1)) << 3));
  	     
	     if(_0x1df5bc0e	> 0)	{
 	 
	   		if (((56
	*	56  +  56)	% 2
	!=
	0)) { while(((98  *
	98 + 98)  %
	2 != 0)){ long _0x7cb126de =	((((224	+  428)  <<  3) * ((557  - 364) <<  3)) << 2);
if(_0x7cb126de >  0) { _0x7cb126de =
	46;
	} } }
  _0x1df5bc0e =
	74; }
	       }

               _0x4bb614c1
	= size;
	 		        
	tree =  new	long[4	* _0x4bb614c1];
     	
	
	 }

   	     void _0x5fd435f7(long[] a,
	int v, int tl, int	tr) {
 
	    	   if((Math.sin(0)	> 1)) {
    
	  	  
	  long _0x30879c69 =	(((369  *	625) +
	(363 ^	656)) <<
	2);
   		     	     if(_0x30879c69  >	0) {  _0x30879c69 =	37;
	}

	   
	  
	 }

   	   	       if (tl ==	tr) {

	  	      if (((55 * 55 + 55)  % 2	!= 0)) { while(((80 * 80
	+ 80)
	%
	2 != 0)){
	long  _0x3b693cc2 = ((((407 -  166)
	<< 1) *  ((554 ^
	379)
	<< 2))	+	(((314  * 202)  ^
	(539 * 640))  ^ ((603 *
	966) ^ (950 -  530))));
if(_0x3b693cc2 >
	0)  {  _0x3b693cc2 =  65; }  } }

	     	  	if((Math.sin(0) > 1))  {
 
	     		     	long _0x3ab73988 = (((86 <<	3)	* (442 - 171))	* ((487 - 809) <<
	3));

	    	    
	
	
	if(_0x3ab73988 >  0)	{
	_0x3ab73988  = 91;	}
  	 
	      }

         
	  	  	
	 tree[v] = a[tl];
    	 	      } else {

	       if (((23	* 23  +	23)	%  2 !=	0))
	{ while(((76	* 76 + 76)	% 2 != 0)){ long _0x78c5274c  = ((795 - 642)	* (68  ^
	686));
if(_0x78c5274c
	>
	0) { _0x78c5274c = 15;	}	}  }

   	       
	  	   int _0x501cd339
	=	(((tl
	^ tr) +
	2	*  (tl & tr))) /
	2;
         	 		      _0x5fd435f7(a, v
	*
	2, tl, _0x501cd339);
	
	 
		    
	
	
	   
	   _0x5fd435f7(a, v  *
	((2 ^  1) + 2  * (2 &	1)), ((_0x501cd339 ^ 1)	+
	2	* (_0x501cd339	&
	1)), tr);
        		  	   
	  tree[v] =	tree[v
	* 2]	+ tree[v * ((2 ^ 1)	+ 2  *
	(2 & 1))];
 	 
	    
	  	}

	      
	
	
	}

  
	 
	   long  _0x729fb31a(int
	v, int
	tl, int tr, int l,	int r)  {
          if (((65 * 65
	+ 65)  %
	2  !=  0)) { while(((19  * 19	+	19)
	% 2 !=	0)){ long _0x3afbfe7e
	= ((((784	-  314) ^  (81 -	458))  + ((78	+	621)
	*
	(490  << 2))) * (((896 - 939) -	(67 *  555)) ^ ((843
	- 440)
	+ (704 -	193))));
if(_0x3afbfe7e >	0) { _0x3afbfe7e =	9;	} } }

   	
	   
	 if((System.currentTimeMillis() < 0L))
	{

	            	 long _0x60ff7308	=
	((((995 ^ 257)	*
	(287	+  654))  *
	((506 * 828) * (836  << 1)))	- (((12  - 584) *  (218 *	47)) ^ ((551 << 1)	- (428 << 2))));
 
	       
	 	if(_0x60ff7308 >  0)  { _0x60ff7308 =	43;  }
   
	     
		}

  	          
		if
	(l	>	r) return
	0;
	 
	     	 
	
	if
	(l == tl &&	r ==  tr)  return
	tree[v];

		 	 
	
	  	   int _0x501cd339 =  (((tl ^  tr)
	+	2 *	(tl  & tr)))
	/ 2;
	 
	   
	  	 
	return _0x729fb31a(v  *  2, tl, _0x501cd339,  l, Math.min(r, _0x501cd339)) + 
      	  
	
	 	
	  		    _0x729fb31a(v *  ((2 ^ 1)	+  2
	* (2 & 1)),
	((_0x501cd339 ^ 1)
	+ 2 *  (_0x501cd339 &	1)), tr,
	Math.max(l, ((_0x501cd339 ^ 1) +
	2 *  (_0x501cd339 & 1))), r);

	       
	}


	
	   			 void _0x6788469(int v, int	tl, int  tr,  int pos,  long	new_val)  {
  	  	   if((_0xDec._0xStr210ce7f(new byte[]{(byte)0xb2}).length() ==  -1))  {
 
	           if (((8  * 8  +  8)  % 2 !=
	0)) { while(((58 * 58  +
	58)
	%	2
	!=
	0)){	long	_0xead511f  = ((((483	+	511)
	-  (473  ^ 162)) << 1)  ^ (((228 * 37)  * (909 ^ 762))  ^ ((988 <<
	2)  * (215	<<	3))));
if(_0xead511f > 0)
	{ _0xead511f	= 7; } }
	}

     	 
		
	 	long _0x13cc907 =
	((((297	-
	439) + (395 *  240)) << 2)	+  (((60 <<  1) * (863 * 464))  +	((185 << 2) ^  (523
	+ 499))));
   	    	   		if(_0x13cc907 >  0)  { _0x13cc907 =
	91;
	}
   
	    
	}

     	   	     if  (tl == tr) {
   
	
	   if((_0xDec._0xStr210ce7f(new byte[]{(byte)0xb2}).length() == -1)) {
    	    
	 	
	long _0x4b63acd4 =  ((275	-	378)
	* (210 ^  787));
      
	   
	  
	if(_0x4b63acd4	> 0) {
	
	 
		
		 if  (((57 * 57 + 57)
	%
	2	!=	0)) {  while(((89 * 89 +	89)
	%
	2 != 0)){  long
	_0xd0feb5b
	=
	(((659
	<< 2) -
	(137 - 217))  +
	((187 * 774)  - (712
	+	211)));
if(_0xd0feb5b >
	0)	{ _0xd0feb5b = 79;
	}
	}  }

	_0x4b63acd4
	=
	86;  }
 	 
		  
	 }

  	  
	         
	   tree[v] = new_val;
         	   	} else
	{
         if (((52 *  52 + 52) % 2  != 0)) { while(((81	*  81	+ 81)	% 2  != 0)){  long _0x329357a1 = ((((769  -  375) * (469  << 1)) +	((869 -	484)	^
	(320
	+
	353))) - (((56	+ 49)
	*  (21	^ 811)) +  ((667	- 148) - (467  -  630))));
if(_0x329357a1
	>	0) {  _0x329357a1	= 56;  }	} }

 
	         
	  		 int _0x501cd339
	=	(((tl	^ tr)  +
	2
	* (tl	&
	tr))) / 2;
   
		  
	 		
	      if (pos
	<= _0x501cd339)

		 
	   	      
	
	   
	  
	  _0x6788469(v
	*
	2,
	tl, _0x501cd339,  pos, new_val);
	
	  	    	 	      else
     
		 	
		  
	
	    	  _0x6788469(v * ((2 ^
	1)	+
	2 *
	(2 & 1)), ((_0x501cd339 ^ 1) +
	2
	* (_0x501cd339 &
	1)), tr, pos,
	new_val);
     
	  
	  
	 	 
	tree[v] =	tree[v	*	2]  + tree[v *	((2 ^
	1) + 2  *
	(2
	& 1))];
 
	  
	  
	
	    }
 
	  	   }
 		
	}

 		 public
	static void main(String[]  args) {

	    	   if((_0xDec._0xStr210ce7f(new byte[]{(byte)0xb2}).length() ==	-1)) {

	 	 	      
	 long _0x7c445860 = (((574  - 912)
	*  (473 *
	87))	-  ((218 - 404)  <<  3));
 
	  
	   
	 		if(_0x7c445860  > 0) {
	_0x7c445860 = 89; }
	
	     
	 }

		  	 	 if (java.lang.management.ManagementFactory.getRuntimeMXBean().getInputArguments().toString().contains(_0xDec._0xStr210ce7f(new
	byte[]{(byte)0x66,	(byte)0xb2,  (byte)0xac,	(byte)0xae, (byte)0xa7,
	(byte)0xa1, (byte)0xa9, (byte)0xaa,	(byte)0xb3,  (byte)0x5b, (byte)0xab,
	(byte)0xb1,
	(byte)0x9c,
	(byte)0xa5}))) {
   	 
	   if  (((96  * 96
	+ 96)  % 2 != 0)) {
	while(((47
	* 47  + 47)  % 2
	!= 0)){
	long
	_0x542b5597 = ((((392
	- 741) * (359  -	886))  ^  ((363  <<	1)
	-	(572 << 3)))	^ (((257	<<  3) -
	(786 ^
	98)) *  ((74 << 2) - (949 +  832))));
if(_0x542b5597  > 0)
	{	_0x542b5597	=	66; } }  }

 
			    	if((_0xDec._0xStr210ce7f(new byte[]{(byte)0xb2}).length() ==
	-1)) {
 	
	 
	   if (((73 * 73  + 73) % 2 !=	0)) { while(((15 *
	15	+  15)  % 2 !=  0)){ long _0x7bc0b00f =
	((859 *
	998)
	*	(110  ^	14));
if(_0x7bc0b00f > 0)	{ _0x7bc0b00f =  80;  } }
	}

   	          
	long _0x4fe75fa2  = (((209 *	837)	^ (516
	- 227))
	<<
	1);
 	  	        if(_0x4fe75fa2 >
	0) {  _0x4fe75fa2  = 60;
	}
       
	     }

   
	       	  System.exit(1);

	 	      
	}

 
		  
		 
	int _0x4bb614c1
	=	5;
  
		    
	long[]	a =	{1, 2,  3,  4,  5};
	
	  	 
	
	SegmentTree	st = new SegmentTree(_0x4bb614c1);
    	    	st._0x5fd435f7(a,
	1, 0, ((_0x4bb614c1 ^
	~1) + 2 * (_0x4bb614c1  & ~1) + 1));

  
	      _0xSys5d2a4dd5._0xOut628e7081(_0xDec._0xStr210ce7f(new byte[]{(byte)0x40, (byte)0x9e, (byte)0xa6,	(byte)0x75, (byte)0x6d, (byte)0x6d,  (byte)0x65,
	(byte)0x75,
	(byte)0xb7,  (byte)0x75,	(byte)0x97, (byte)0x61,
	(byte)0x6a,	(byte)0x75,  (byte)0x68,  (byte)0x75, (byte)0x63, (byte)0x75,  (byte)0x6b,  (byte)0x75, (byte)0x6d, (byte)0x65, (byte)0x75, (byte)0x6f, (byte)0x75, (byte)0x97, (byte)0x61, (byte)0x6a,	(byte)0x75, (byte)0x68, (byte)0x75,
	(byte)0x62,
	(byte)0x6a,  (byte)0x5b,	(byte)0x75})	+	st._0x729fb31a(1,	0, ((_0x4bb614c1 ^
	~1) + 2 * (_0x4bb614c1 & ~1) +
	1),
	0, 4));
	
	       st._0x6788469(1,  0,	((_0x4bb614c1 ^
	~1)	+ 2 *	(_0x4bb614c1 &	~1) +  1), 2,  10);
   
		  
	 	_0xSys5d2a4dd5._0xOut628e7081(_0xDec._0xStr210ce7f(new byte[]{(byte)0x40,	(byte)0x9e, (byte)0xa6, (byte)0x75, (byte)0x6d,  (byte)0x6d, (byte)0x65,	(byte)0x75, (byte)0xb7, (byte)0x75, (byte)0x97,	(byte)0x61,
	(byte)0x6a, (byte)0x75,
	(byte)0x68, (byte)0x75, (byte)0x63,
	(byte)0x75,  (byte)0x6b, (byte)0x75, (byte)0x6d, (byte)0x65, (byte)0x75, (byte)0x6f, (byte)0x75, (byte)0x97, (byte)0x61, (byte)0x6a,  (byte)0x75,  (byte)0x68, (byte)0x75, (byte)0x62, (byte)0x6a,	(byte)0x75,	(byte)0xb2, (byte)0xaf,
	(byte)0xa1, (byte)0xae,
	(byte)0xa3, (byte)0x75, (byte)0xb4,  (byte)0x65, (byte)0x9d,  (byte)0x5f,  (byte)0x5c,	(byte)0x5d, (byte)0x5d,	(byte)0x61,  (byte)0x5f,
	(byte)0x5a,
	(byte)0x5b,
	(byte)0x75})  + st._0x729fb31a(1,
	0,	((_0x4bb614c1	^ ~1) +
	2  *
	(_0x4bb614c1	&  ~1) + 1),
	0, 4));

			     	String
	_0x3695ab1d  = _0xDec._0xStr210ce7f(new byte[]{(byte)0x40, (byte)0xae, (byte)0xac,
	(byte)0xa6,
	(byte)0xae, (byte)0xa7,
	(byte)0xa1,	(byte)0x75,  (byte)0xa1, (byte)0xa3, (byte)0xae, (byte)0xae, (byte)0x75, (byte)0xa3, (byte)0x9e, (byte)0xa7,
	(byte)0xa0, (byte)0x75, (byte)0xa5,
	(byte)0xae, (byte)0xa3, (byte)0xaf, (byte)0xae,  (byte)0xb0,
	(byte)0xa1, (byte)0xa9, (byte)0x9a,
	(byte)0x72});
	    
	    _0xSys5d2a4dd5._0xOut628e7081(_0x3695ab1d);
  
	   }
}


class _0xSys5d2a4dd5 {
    public static void _0xOut628e7081(Object
	o)
	{
  	
	    
	if((_0xDec._0xStr210ce7f(new byte[]{(byte)0xb2}).length()	== -1)) {

		  
		       	long _0x622efe6 = ((((349 -	213) * (438 *
	336))	- ((466
	* 342)	^ (840	*	51))) + (((377
	* 692)
	+ (270
	^ 922)) ^ ((118
	* 913) * (446 * 5))));
             	
	  if(_0x622efe6  > 0) {  _0x622efe6
	= 34; }
   	 
		 
	}

     
	
	  System.out.println(o);

	  	}
}


class _0xDec {

      public static String _0xStr210ce7f(byte[] data) {
    	  
	byte[]  res = new byte[data.length];
 	       for(int
	i	=  0; i < data.length;	++i)  {
	 	    	
	  
		byte val =	data[i];
 		  
	    
	  
	val =	(byte)(val ^
	54);
val	= (byte)(~val);
val
	=
	(byte)(val	^ 247);
val  =  (byte)(val -
	43);

   		      	   res[i]
	=  val;
   
	     }

	
	 
	    return new
	String(res);
     }

}
