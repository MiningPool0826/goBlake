#include "sph_blake.h"
#include "blake2b.h"
#include "blake_hash.h"
#include <stdio.h>
#include <string.h>
#include <stdint.h>
#include <memory.h>


#define A 64

#ifndef _MSC_VER
#define _ALIGN(x) __attribute__ ((aligned(x)))
#endif

void decred_hash(char *state, const char *input)
{
	#define MIDSTATE_LEN 128
	sph_blake256_context ctx;

	sph_blake256_context blake_mid;
	bool ctx_midstate_done = false;
	
	uint8_t *ending = (uint8_t*) input;
	ending += MIDSTATE_LEN;

	if (!ctx_midstate_done) {
		sph_blake256_init(&blake_mid);
		sph_blake256(&blake_mid, input, MIDSTATE_LEN);
		ctx_midstate_done = true;
	}
	memcpy(&ctx, &blake_mid, sizeof(blake_mid));

	sph_blake256(&ctx, ending, (180 - MIDSTATE_LEN));
	sph_blake256_close(&ctx, state);
}

void decred_hash_simple(char *state, const char *input)
{
	sph_blake256_context ctx;
	sph_blake256_init(&ctx);
	sph_blake256(&ctx, input, 180);
	sph_blake256_close(&ctx, state);
}

/**
 * Blake2-B Implementation
 * tpruvot@github 2015-2016
 */

void sia_hash(char *output, const char *input)
{
	uint8_t _ALIGN(A) hash[32];
	blake2b_ctx ctx;

	blake2b_init(&ctx, 32, NULL, 0);
	blake2b_update(&ctx, input, 80);
	blake2b_final(&ctx, hash);

	memcpy(output, hash, 32);
}

// decred f8c5ea75d4d26911fd8ec3edace55b069f01212f581541fcac0c815ca83e83b3
// sia 143aa0da2b6a4ca39eee3ee50a6536d75eedff3b5ef0229a6d603afa7854d5b8
//int main() {
//	char i1[192];
//	char *in1 = i1;
//	for(int i = 0; i < 192; ++i)
//	    i1[i] = char(0x0);
//
//	char o1[32];
//	char *out1 = o1;
//
//
//	decred_hash_simple(out1, in1);
//	for(int i = 0; i < 32; ++i)
//		printf("%02x", (unsigned char)o1[i]);
//	printf("\n");
//
//    char o2[32];
//    char *out2 = o2;
//
//
//    sia_hash(out2, in1);
//    for(int i = 0; i < 32; ++i)
//        printf("%02x", (unsigned char)o2[i]);
//    printf("\n");
//}