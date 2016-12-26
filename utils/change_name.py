import glob
import os

root_path = "../junior/"

for file in glob.glob(os.path.join(root_path, "problem_*")):
    # print(file)
    # print(file.split("problem_")[1])
    os.rename(file, os.path.join(root_path, "p" + file.split("problem_")[1]))