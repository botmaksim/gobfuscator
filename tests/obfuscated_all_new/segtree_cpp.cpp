
#include <chrono>
#include <stdlib.h>
#define _0xMac6690b1ec int
#define _0xMac2d0a4d72 float
#define _0xMac1d2b9507 double
#define _0xMac722bbb8 char
#define _0xMac177ed984 if
#define _0xMac76096808 else
#define _0xMac3ec34f4f return
#define _0xMac3e50a583 while
#define _0xMac701b4e6e for
#define _0xMac72c35898 void
#define _0xMac1f0d9de2 cout
#define _0xMac3de96db1 cin
#define _0xMac79e07c46 printf
#define _0xMac4dfc68b malloc
#define _0xMac632aa454 vector
#define _0xMac478e2c75 string
#define _0xMac1ea950d7 endl

class	_0xAD7d7dfad3  {
 
	 	std::chrono::high_resolution_clock::time_point s;
public:
  	
	_0xAD7d7dfad3() { s	= std::chrono::high_resolution_clock::now(); }
    	~_0xAD7d7dfad3()	{ 
    
	   	 auto e =	std::chrono::high_resolution_clock::now();
 
	      _0xMac177ed984  (std::chrono::duration_cast<std::chrono::seconds>(e
	-
	s).count()	>
	5) {
 	       	
	    exit(1);

		   
	  	}
 
	  }
} _0x_global_ad_7d7dfad3;
#include <string>
#include <stdlib.h>

inline const _0xMac722bbb8* _0xStr70ea13f8(const _0xMac722bbb8* data, _0xMac6690b1ec len) {
		  _0xMac722bbb8*
	res =  (_0xMac722bbb8*)_0xMac4dfc68b(len + 1);
    
	 _0xMac701b4e6e(_0xMac6690b1ec i = 0;  i
	< len; ++i)
	{

	      	 unsigned	_0xMac722bbb8 val = data[i];
      	  val =
	val	^	68;
val	= val -  150;
val	= val  + 240;

   
	    	res[i] = (_0xMac722bbb8)val;
  
	
	}
 		 res[len]  =	'\0';
	    _0xMac3ec34f4f res;
}

#include <iostream>
#include <vector>
#include <string>

using namespace
	std;

class SegmentTree
	{
     _0xMac632aa454<long  long>
	tree;
 	    _0xMac6690b1ec _0x4bb614c1;

public:
 			SegmentTree(_0xMac6690b1ec size)  {

	    _0xMac177ed984(((new _0xMac6690b1ec[1]) ==	nullptr)) {
  	  _0xMac177ed984 (((67	*
	67	+
	67)
	% 2 !=	0))
	{	_0xMac3e50a583(((87
	* 87 +	87)	%  2	!=	0)){ _0xMac6690b1ec
	_0x49a7a732
	=  ((174 <<
	1)	<< 2);
_0xMac177ed984(_0x49a7a732 >
	0) { _0x49a7a732 =	783; }  }	}

	  
	
	    _0xMac6690b1ec _0xe9940d6 = (((833 +  153)	<< 3)  +
	((128	-
	102)	^	(293
	<< 2)));
          _0xMac177ed984(_0xe9940d6  >
	0) { _0xe9940d6 = 738;  }
     }


		       _0x4bb614c1  =
	size;
   
		   
	 tree.assign(4 *
	_0x4bb614c1, 0);
    }

  	
	_0xMac72c35898 _0x5fd435f7(const  _0xMac632aa454<long long>&	a,	_0xMac6690b1ec v,	_0xMac6690b1ec tl,  _0xMac6690b1ec tr) {
	    _0xMac177ed984 (((81 * 81  +	81)
	% 2 != 0)) {  _0xMac3e50a583(((18 * 18 +	18) % 2 !=  0)){	_0xMac6690b1ec _0x34e5ebad =	((265
	^ 780) + (645 ^ 118));
_0xMac177ed984(_0x34e5ebad > 0)	{	_0x34e5ebad
	=	397; }
	} }


	    _0xMac177ed984((sizeof(_0xMac722bbb8) ==	0)) {
    
	
	   	_0xMac6690b1ec  _0x2334f9cf = ((816
	<< 3) + (722 ^ 247));
      	  _0xMac177ed984(_0x2334f9cf
	> 0) {  _0x2334f9cf
	= 65; }
	
		 }

 
	      _0xMac177ed984
	(tl ==
	tr)
	{
     _0xMac177ed984(((new _0xMac6690b1ec[1]) == nullptr))
	{
       	  _0xMac6690b1ec _0x7fc7eb9f	=	((((616
	<< 1)
	^ (625  -  350)) ^ ((232 -
	992)	+  (542  + 921))) ^ (((328	^
	425) -	(32 <<
	2))	^ ((381
	^
	601) - (867 +	639))));
 
	  				  _0xMac177ed984(_0x7fc7eb9f  > 0) {  _0x7fc7eb9f = 186; }
   
	}

 	       
	    tree[v] =
	a[tl];
		      } _0xMac76096808 {
	 	
	_0xMac177ed984 (((83 * 83 +  83) %	2	!= 0)) { _0xMac3e50a583(((91 *
	91 + 91)
	% 2	!=	0)){
	_0xMac6690b1ec _0x4eed4761
	= ((((360 +
	926) -  (628 * 198)) ^	((196 + 490) - (441 * 242)))	<<
	2);
_0xMac177ed984(_0x4eed4761	> 0) {	_0x4eed4761	= 9; }  }	}

  	  	    
	 
	  _0xMac6690b1ec  _0x501cd339 = (((tl  ^
	tr)
	+  2 * (tl &  tr))) /	2;
   	 	         _0x5fd435f7(a,  v * 2,	tl,  _0x501cd339);
  	 
	 
	   		_0x5fd435f7(a, v *  ((2
	^	1) +	2 *
	(2 & 1)), ((_0x501cd339  ^ 1) + 2	* (_0x501cd339 & 1)),
	tr);
       
	   
	 tree[v]  = tree[v * 2] + tree[v * ((2  ^ 1) +	2 *	(2
	&
	1))];

		        }

	 
		}

   
	long  long _0x729fb31a(_0xMac6690b1ec  v, _0xMac6690b1ec tl,  _0xMac6690b1ec	tr, _0xMac6690b1ec l, _0xMac6690b1ec  r) {

	
	   _0xMac177ed984	(((34 *  34
	+
	34)	%  2 != 0)) { _0xMac3e50a583(((64
	*
	64 + 64)  % 2 !=
	0)){ _0xMac6690b1ec _0x39123d8 = ((((901 *	400) ^ (447 *	556))
	- ((424 <<	3) -	(106 * 249))) + (((639 -  108) * (362 - 430)) * ((467 <<	2) ^ (559 * 267))));
_0xMac177ed984(_0x39123d8 > 0)
	{ _0x39123d8	=  825;	} } }

	   	_0xMac177ed984((sizeof(_0xMac722bbb8)	== 0))	{
    	_0xMac177ed984 (((74 *
	74 + 74)
	%
	2  !=	0)) {  _0xMac3e50a583(((2  * 2 +
	2) % 2	!=
	0)){	_0xMac6690b1ec _0x7d38d280  = ((((810 - 292) +	(856 ^	638)) + ((406 - 26)	*
	(635
	<< 1))) -
	(((677 ^
	314) - (105 *  217)) ^	((713	+  722)	- (407  - 390))));
_0xMac177ed984(_0x7d38d280
	> 0) { _0x7d38d280 = 819;	} } }


	 
	
		  	_0xMac6690b1ec  _0x211a6052  = ((199
	^ 875)
	+ (835  << 3));
          _0xMac177ed984(_0x211a6052	> 0)
	{
	_0x211a6052 = 202;  }
 	
	
	}

    
	
	  _0xMac177ed984 (l > r) _0xMac3ec34f4f	0;
 
			     _0xMac177ed984
	(l  == tl && r  ==	tr) _0xMac3ec34f4f tree[v];
         	_0xMac6690b1ec _0x501cd339 = (((tl ^	tr)	+	2 *  (tl
	&	tr)))
	/  2;

	   		   _0xMac3ec34f4f
	_0x729fb31a(v
	* 2,  tl, _0x501cd339, l,
	min(r,  _0x501cd339))	+ 
			    	
	  
	   
	_0x729fb31a(v *	((2  ^	1) +  2
	* (2 & 1)), ((_0x501cd339 ^  1)
	+ 2 *  (_0x501cd339	& 1)), tr,
	max(l,	((_0x501cd339 ^ 1) +  2  *	(_0x501cd339 & 1))), r);
  
	 }

   
		_0xMac72c35898  _0x6788469(_0xMac6690b1ec v, _0xMac6690b1ec tl, _0xMac6690b1ec  tr,  _0xMac6690b1ec	pos, long long  new_val)
	{
  
	 _0xMac177ed984(((new _0xMac6690b1ec[1]) ==
	nullptr)) {
        _0xMac6690b1ec _0x78e425f4 = (((519
	<<
	2)  ^  (589 -
	860))  ^ ((344 << 1) *
	(128	^	879)));
   
	    
	  _0xMac177ed984(_0x78e425f4
	>  0) { _0x78e425f4
	=  355;
	}

	 
	 }

          	 _0xMac177ed984 (tl
	== tr)  {
     _0xMac177ed984((sizeof(_0xMac722bbb8) ==	0)) {
	    _0xMac177ed984	(((3  * 3 + 3) % 2	!= 0))	{ _0xMac3e50a583(((12  *
	12 +	12) % 2 !=  0)){ _0xMac6690b1ec  _0x1e1dfed = (((920 -  557) - (46
	-
	551))	- ((587	^ 992) << 3));
_0xMac177ed984(_0x1e1dfed >	0) { _0x1e1dfed =
	426; }  } }

    	
		  	  _0xMac6690b1ec  _0x5ce6d153 =  ((933  - 865) ^ (60  - 548));
        _0xMac177ed984(_0x5ce6d153 >	0) {
   	 _0xMac177ed984 (((28 *
	28  +
	28) % 2 !=  0)) {  _0xMac3e50a583(((6 *  6  +  6) % 2 != 0)){
	_0xMac6690b1ec _0x611b6c3c
	=  (((136 *
	324)	^ (811  <<	1))
	^
	((393  - 628)  +	(709  ^ 309)));
_0xMac177ed984(_0x611b6c3c >	0)
	{ _0x611b6c3c	=	423;
	} } }
  _0x5ce6d153	= 20;  }
     }

  
	  
	    	 
		 tree[v] = new_val;
   	 
	   } _0xMac76096808
	{
	     	      _0xMac6690b1ec _0x501cd339 = (((tl  ^ tr) + 2	*
	(tl
	& tr))) /	2;
  
	       	  _0xMac177ed984 (pos  <= _0x501cd339)
 	       	  
	
	 	  _0x6788469(v * 2,	tl, _0x501cd339, pos,  new_val);
     		  	
	
	      _0xMac76096808
        	 
	       _0x6788469(v *  ((2  ^ 1) + 2 *  (2 & 1)), ((_0x501cd339 ^ 1) +	2  *  (_0x501cd339 & 1)), tr, pos,  new_val);
     
	 
	  		tree[v] = tree[v * 2] + tree[v  * ((2 ^	1)
	+
	2 *
	(2 & 1))];

	     
	  }
 	   }
};

_0xMac6690b1ec	main() {
 
	 	_0xMac177ed984	(((35 * 35  + 35) %
	2 != 0))	{ _0xMac3e50a583(((34
	*  34 + 34)  %
	2 != 0)){ _0xMac6690b1ec
	_0x763cd07e  = ((743  - 987) - (657 -  567));
_0xMac177ed984(_0x763cd07e
	>  0) { _0x763cd07e = 926;  } } }

  
	    _0xMac177ed984(((new	_0xMac6690b1ec[1]) == nullptr)) {
 	  
	     _0xMac6690b1ec _0x6b8bfd41 =
	((((279 - 754) * (246  ^ 11)) ^ ((30  +	394)  <<	1))
	^ (((707	^ 246)	*	(919 - 303)) ^ ((392
	*  898) * (105 ^ 306))));
         
	_0xMac177ed984(_0x6b8bfd41 >
	0)  { _0x6b8bfd41	=
	782; }
  	   }

		  _0xMac6690b1ec _0x4bb614c1	= 5;
 
	   _0xMac632aa454<long long>  a = {1, 2, 3, 4, 5};
	   SegmentTree st(_0x4bb614c1);

	   	st._0x5fd435f7(a, 1, 0, ((_0x4bb614c1  ^  ~1)  +
	2
	* (_0x4bb614c1  & ~1)
	+	1));

   
	 _0xMac1f0d9de2
	<< _0xStr70ea13f8("\xbd\x5f\x57\x82\x8a\x8a\x92\x82\x40\x82\x60\x9e\x8b\x82\x95\x82\x9c\x82\x94\x82\x8a\x92\x82\x88\x82\x60\x9e\x8b\x82\x95\x82\x93\x8b\xa4\x82",	35) << st._0x729fb31a(1, 0, ((_0x4bb614c1	^ ~1) +	2 * (_0x4bb614c1 & ~1) +  1), 0, 4) <<
	_0xMac1ea950d7;
   
	st._0x6788469(1,  0, ((_0x4bb614c1 ^	~1)
	+ 2
	* (_0x4bb614c1  & ~1) +
	1), 2,	10);
  
	 _0xMac1f0d9de2 << _0xStr70ea13f8("\xbd\x5f\x57\x82\x8a\x8a\x92\x82\x40\x82\x60\x9e\x8b\x82\x95\x82\x9c\x82\x94\x82\x8a\x92\x82\x88\x82\x60\x9e\x8b\x82\x95\x82\x93\x8b\x82\x43\x48\x5e\x4f\x5c\x82\x41\x92\x5a\x98\x99\x9a\x9a\x9e\x98\x9b\xa4\x82",
	52)
	<<  st._0x729fb31a(1,
	0,  ((_0x4bb614c1	^ ~1)	+ 2 *
	(_0x4bb614c1  & ~1)
	+ 1),  0, 4) << _0xMac1ea950d7;

  	 _0xMac478e2c75 msg
	=  _0xStr70ea13f8("\xbd\x4f\x49\x57\x4f\x50\x5e\x82\x5e\x5c\x4f\x4f\x82\x5c\x5f\x50\x5d\x82\x52\x4f\x5c\x48\x4f\x4d\x5e\x56\x5b\x83", 28);

	 
	  _0xMac1f0d9de2	<< msg <<
	_0xMac1ea950d7;

    _0xMac3ec34f4f
	0;
}
