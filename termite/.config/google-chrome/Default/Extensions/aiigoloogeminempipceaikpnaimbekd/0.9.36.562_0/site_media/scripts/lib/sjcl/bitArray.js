/** @fileOverview Arrays of bits, encoded as arrays of Numbers.
 *
 * @author Emily Stark
 * @author Mike Hamburg
 * @author Dan Boneh
 */

/** @namespace Arrays of bits, encoded as arrays of Numbers.
 *
 * @description
 * <p>
 * These objects are the currency accepted by SJCL's crypto functions.
 * </p>
 *
 * <p>
 * Most of our crypto primitives operate on arrays of 4-byte words internally,
 * but many of them can take arguments that are not a multiple of 4 bytes.
 * This library encodes arrays of bits (whose size need not be a multiple of 8
 * bits) as arrays of 32-bit words.  The bits are packed, big-endian, into an
 * array of words, 32 bits at a time.  Since the words are double-precision
 * floating point numbers, they fit some extra data.  We use this (in a private,
 * possibly-changing manner) to encode the number of bits actually  present
 * in the last word of the array.
 * </p>
 *
 * <p>
 * Because bitwise ops clear this out-of-band data, these arrays can be passed
 * to ciphers like AES which want arrays of words.
 * </p>
 */
sjcl.bitArray={bitSlice:function(a,bstart,bend){a=sjcl.bitArray._shiftRight(a.slice(bstart/32),32-(bstart&31)).slice(1);return(bend===undefined)?a:sjcl.bitArray.clamp(a,bend-bstart);},extract:function(a,bstart,blength){var x,sh=Math.floor((-bstart-blength)&31);if((bstart+blength-1^bstart)&-32){x=(a[bstart/32|0]<<(32-sh))^(a[bstart/32+1|0]>>>sh);}else{x=a[bstart/32|0]>>>sh;}
return x&((1<<blength)-1);},concat:function(a1,a2){if(a1.length===0||a2.length===0){return a1.concat(a2);}
var out,i,last=a1[a1.length-1],shift=sjcl.bitArray.getPartial(last);if(shift===32){return a1.concat(a2);}else{return sjcl.bitArray._shiftRight(a2,shift,last|0,a1.slice(0,a1.length-1));}},bitLength:function(a){var l=a.length,x;if(l===0){return 0;}
x=a[l-1];return(l-1)*32+sjcl.bitArray.getPartial(x);},clamp:function(a,len){if(a.length*32<len){return a;}
a=a.slice(0,Math.ceil(len/32));var l=a.length;len=len&31;if(l>0&&len){a[l-1]=sjcl.bitArray.partial(len,a[l-1]&0x80000000>>(len-1),1);}
return a;},partial:function(len,x,_end){if(len===32){return x;}
return(_end?x|0:x<<(32-len))+len*0x10000000000;},getPartial:function(x){return Math.round(x/0x10000000000)||32;},equal:function(a,b){if(sjcl.bitArray.bitLength(a)!==sjcl.bitArray.bitLength(b)){return false;}
var x=0,i;for(i=0;i<a.length;i++){x|=a[i]^b[i];}
return(x===0);},_shiftRight:function(a,shift,carry,out){var i,last2=0,shift2;if(out===undefined){out=[];}
for(;shift>=32;shift-=32){out.push(carry);carry=0;}
if(shift===0){return out.concat(a);}
for(i=0;i<a.length;i++){out.push(carry|a[i]>>>shift);carry=a[i]<<(32-shift);}
last2=a.length?a[a.length-1]:0;shift2=sjcl.bitArray.getPartial(last2);out.push(sjcl.bitArray.partial(shift+shift2&31,(shift+shift2>32)?carry:out.pop(),1));return out;},_xor4:function(x,y){return[x[0]^y[0],x[1]^y[1],x[2]^y[2],x[3]^y[3]];}};