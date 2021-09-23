# Recursive Subdirectory Removal tool

The intent of this program is to be able to remove lower level subdirectories
from a parent path.
In Unity projects (when backing up, etc.), the Library
folder, in particular can grow gigabytes large. This folder
is easily recreated when reopening the project and isn't needed
when backing up a project. 

This program provides a way to go through a parent directory 
of projects and remove these directories.

e.g. rsr.exe D:\UnityProjects Library

provides a non-destructive print out of what would be deleted.
Including a "Y" command line parm at the end tells the program
to actually delete the files..

e.g. rsr.exe D:\UnityProjects Library y