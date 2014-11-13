/** @fileOverview HMAC implementation.
 *
 * @author Emily Stark
 * @author Mike Hamburg
 * @author Dan Boneh
 */

/** HMAC with the specified hash function.
 * @constructor
 * @param {bitArray} key the key for HMAC.
 * @param {Object} [hash=sjcl.hash.sha256] The hash function to use.
 */
sjcl.misc.hmac=function(key,Hash){this._hash=Hash=Hash||sjcl.hash.sha256;var exKey=[[],[]],i,bs=Hash.prototype.blockSize/32;this._baseHash=[new Hash(),new Hash()];if(key.length>bs){key=Hash.hash(key);}
for(i=0;i<bs;i++){exKey[0][i]=key[i]^0x36363636;exKey[1][i]=key[i]^0x5C5C5C5C;}
this._baseHash[0].update(exKey[0]);this._baseHash[1].update(exKey[1]);this._resultHash=new Hash(this._baseHash[0]);};sjcl.misc.hmac.prototype.encrypt=sjcl.misc.hmac.prototype.mac=function(data){this.update(data);return this.digest(data);};sjcl.misc.hmac.prototype.update=function(data){this._resultHash.update(data);};sjcl.misc.hmac.prototype.digest=function(){var w=this._resultHash.finalize();return new(this._hash)(this._baseHash[1]).update(w).finalize();};