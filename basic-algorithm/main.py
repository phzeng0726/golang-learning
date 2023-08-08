# %%

strs = ["flower","flow","flight"]

i=1
isSame = True
while isSame:
    if len(set([str[0:i] for str in strs])) == 1:
        i+=1
    else:
        isSame = False
print(strs[0][0:i])


#%%
