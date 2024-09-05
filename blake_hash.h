#ifndef __BLAKEHASH_H__
#define __BLAKEHASH_H__

#ifdef __cplusplus
extern "C" {
#endif

void decred_hash(char *state, const char *input);
void decred_hash_simple(char *state, const char *input);
void sia_hash(char *output, const char *input);

#ifdef __cplusplus
}
#endif

#endif
