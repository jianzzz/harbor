import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';

import { SharedModule } from '../shared/shared.module';

import { RepositoryComponent } from './repository.component';
import { ListRepositoryComponent } from './list-repository/list-repository.component';
import { TagRepositoryComponent } from './tag-repository/tag-repository.component';

import { RepositoryService } from './repository.service';

@NgModule({
  imports: [
    SharedModule,
    RouterModule
  ],
  declarations: [
    RepositoryComponent,
    ListRepositoryComponent,
    TagRepositoryComponent
  ],
  exports: [RepositoryComponent, ListRepositoryComponent],
  providers: [RepositoryService]
})
export class RepositoryModule { }